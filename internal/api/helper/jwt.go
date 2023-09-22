package helper

import (
	"errors"
	"time"

	// "github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/db"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
)

type UserAccessTokenClaims struct {
	User      uuid.UUID         `json:"user"`
	Payload   map[string]string `json:"payload"`
	TokenType string            `json:"type"`
	Role      *uuid.UUID        `json:"role"`
	jwt.StandardClaims
}

type UserRefreshTokenClaims struct {
	User      uuid.UUID         `json:"user"`
	Payload   map[string]string `json:"payload"`
	TokenType string            `json:"type"`
	jwt.StandardClaims
}

// Generate access and refresh tokens
func GenerateJwtTokens(userId uuid.UUID, roleId *uuid.UUID) (*models.Token, error) {
	env := db.NewEnv()

	accessTokenKey := env.AccessTokenSecret
	refreshTokenKey := env.RefreshTokenSecret

	//time for refresh token in hour
	var refreshTimeAdded = env.RefreshTokenExpiryHour

	//time for access token in hour
	var accessTimeAdded = env.AccessTokenExpiryHour

	issuedAt := time.Now().Unix()
	//access token claims
	accessTokenClaims := UserAccessTokenClaims{
		User:      userId,
		Role:      roleId,
		Payload:   map[string]string{},
		TokenType: "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(accessTimeAdded)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	//refresh token claims
	refreshTokenClaims := UserRefreshTokenClaims{
		User:      userId,
		Payload:   map[string]string{},
		TokenType: "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(refreshTimeAdded)).Unix(),
			IssuedAt:  issuedAt,
		},
	}

	//refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	//access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	//get a signed access token string
	accessTokenString, accessError := accessToken.SignedString([]byte(accessTokenKey))
	// Sign and get the complete encoded token as a string using the secret
	refreshTokenString, refreshError := refreshToken.SignedString([]byte(refreshTokenKey))

	if accessError != nil || refreshError != nil {
		return nil, errors.New("error generating token")
	}
	return &models.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		Iat:          issuedAt,
	}, nil

}

func ParseToken(database *gorm.DB, tokenStr string) (*uuid.UUID, error) {
	//get secret key
	env := db.NewEnv()
	secretKey := []byte(env.RefreshTokenSecret)

	token, err := jwt.ParseWithClaims(tokenStr, &UserRefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error with signature check")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	//get user id
	var userId uuid.UUID
	var iat int64
	if claims, ok := token.Claims.(*UserRefreshTokenClaims); ok && token.Valid {
		userId = claims.User
		iat = claims.IssuedAt
	} else {
		return nil, err
	}

	var user models.User
	if err := database.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	lastRevokedAt := user.TokenLastRevokedAt
	issuedAt := time.Unix(iat, 0)
	if issuedAt.Before(lastRevokedAt) {
		return nil, errors.New("invalid token")
	}

	return utils.Ref(userId), nil
}
