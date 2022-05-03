package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenUserData struct {
	Issuer     string
	IssuedAt   string
	Expiration string
	UserID     string
}

type JWTLoginToken struct {
	UserData TokenUserData `json:"userData"`
	jwt.RegisteredClaims
}
