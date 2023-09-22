package api

import (
	"encoding/base64"

	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashData(data string) (string, error) {
	var bytes []byte
	var err error
	if data != "" {
		bytes, err = bcrypt.GenerateFromPassword([]byte(data), 14)
	}

	return string(bytes), err
}

func GenerateSecureToken(id uuid.UUID) (string, error) {
	uuidBytes, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}

	// Encode the resulting byte slice to base64 string
	token := base64.URLEncoding.EncodeToString(uuidBytes)

	return token, nil
}

func GetDataFromSecureToken(token string) (*uuid.UUID, error) {
	uuidBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return nil, err
	}

	return utils.Ref(userId), nil
}

const LinkExpireInHours = 2

type TableName string

const (
	Invitations TableName = "invitation"
	Bride       TableName = "bride"
	Groom       TableName = "groom"
)
