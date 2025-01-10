package auth

import (
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(customClaim models.CustomClaims) (string, error)
	// ValidateToken(token string) (string, error)
}

type jwtService struct {
	config *config.Configuration
}

func NewJWTService(config *config.Configuration) JWTService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateToken(customClaim models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim.Register)

	tokenString, err := token.SignedString([]byte(s.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
