package models

import "time"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Iat          int64  `json:"-"`
}

type UserLoginView struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenView struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserLoginAttempt struct {
	FailedAttempt  int       `json:"-"`
	LockedUntil    time.Time `json:"-"`
	LockedDuration int       `json:"-"`
}
