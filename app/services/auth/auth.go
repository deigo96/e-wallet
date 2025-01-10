package auth

import (
	"context"
	"errors"
	"time"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/roles"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(c context.Context, request *models.LoginRequest) (*models.LoginResponse, error)
}

type authService struct {
	userRepository users.UserRepository
	roleRepository roles.RoleRepository
	jwtService     JWTService
	config         *config.Configuration
}

func NewAuthService(config *config.Configuration, db *gorm.DB) AuthService {
	return &authService{
		userRepository: users.NewUserRepository(db),
		roleRepository: roles.NewRoleRepository(db),
		jwtService:     NewJWTService(config),
		config:         config,
	}
}

func (as *authService) Login(c context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	user, err := as.userRepository.GetUserByEmail(c, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrIncorrectEmailOrPassword
		}
		return nil, err
	}

	if !as.isVerifiedPassword(request.Password, user.Password) {
		return nil, customError.ErrIncorrectEmailOrPassword
	}

	token, err := as.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:     token,
		TokenType: constant.TokenTypeBearer,
	}, nil
}

func (as *authService) isVerifiedPassword(password, hashedPassword string) bool {
	return utils.ValidatePassword(password, hashedPassword)
}

func (as *authService) generateToken(user *entity.User) (string, error) {
	customClaim := models.CustomClaims{
		Email: user.Email,
		ID:    user.ID,
		Role:  constant.GetRoleName(user.Role),
		Register: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    as.config.ServiceName,
		},
	}

	token, err := as.jwtService.GenerateToken(customClaim)
	if err != nil {
		return "", err
	}

	return token, nil
}
