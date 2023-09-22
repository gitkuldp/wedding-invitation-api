package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gitkuldp/wedding-invitation-api/internal/api/helper"
	"github.com/gitkuldp/wedding-invitation-api/internal/db"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Authenticator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		user := models.User{
			ID:    utils.StringToUUID("f3e9a178-337b-40e3-ab35-fa581d41d1e5"),
			Email: "lamichhane.kebib@gmail.com",
			Phone: "9779806791681",
		}
		e.Set("user", user)
		return next(e)
	}
}

func Authentication(database *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			//get the token from request
			authenticationToken := e.Request().Header.Get("Authorization")
			if authenticationToken == "" || len(authenticationToken) < 7 {
				return echo.NewHTTPError(http.StatusBadRequest, "missing token")
			}

			authenticationTokenSplit := strings.Split(authenticationToken, " ")

			if len(authenticationTokenSplit) != 2 {
				return echo.NewHTTPError(http.StatusBadRequest, "missing token")
			}

			authenticationToken = authenticationTokenSplit[1]

			//get secret key
			env := db.NewEnv()
			secretKey := []byte(env.AccessTokenSecret)

			token, err := jwt.ParseWithClaims(authenticationToken, &helper.UserAccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("error with signature check")
				}

				return secretKey, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token  1")
			}

			var userId uuid.UUID
			var roleId *uuid.UUID
			if claims, ok := token.Claims.(*helper.UserAccessTokenClaims); ok && token.Valid {
				userId = claims.User
				roleId = claims.Role
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			//check if user is active or not and user is associated with the role
			var user models.User
			if err := database.Joins("LEFT JOIN user_roles ON users.id = user_roles.user_id").
				First(&user, "users.id = ? AND NOT users.is_active AND (user_roles.role_id = ? OR user_roles.role_id IS NULL)", userId, roleId).Error; err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid user")
			}

			//Get permissions
			var roleReference models.RoleReference
			if roleId != nil {
				if err := database.Where("id = ?", roleId).First(&roleReference).Error; err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "invalid role")
				}
			}

			e.Set("permissions", roleReference.Permissions)
			e.Set("user", user)
			return next(e)
		}
	}
}
