package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//go:generate easyjson types.go

// CustomClaims for store merchant_id
// easyjson:json
type CustomClaims struct {
	jwt.RegisteredClaims
	MerchantUUID uuid.UUID `json:"merchant_id"`
}
