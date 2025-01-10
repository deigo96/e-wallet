package models

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type CustomClaims struct {
	ID       int
	Role     string
	Email    string
	Register jwt.RegisteredClaims
}
