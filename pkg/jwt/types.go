package jwtvalitator

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//go:generate easyjson types.go
// easyjson:json

type CustomClaims struct {
	jwt.RegisteredClaims
	MerchantUUID uuid.UUID `json:"merchant_id"`
	WalletUUID uuid.UUID `json:"wallet_id"`
}
