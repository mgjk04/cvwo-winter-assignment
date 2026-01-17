package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

//DTO
type LoginReq struct {
	Username string `json:"username" binding:"required,excludesall= "`
}

//same as login but I'll keep here for future extnsibility
type SignupReq struct {
	Username string `json:"username" binding:"required,excludesall= "`
}