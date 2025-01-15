package auth

import (
	"context"
	"errors"
	"time"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/external"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/otp"
	"github.com/deigo96/e-wallet.git/app/repository/profile"
	"github.com/deigo96/e-wallet.git/app/repository/roles"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(c context.Context, request *models.LoginRequest) (*models.LoginResponse, error)
	SendOTP(c *gin.Context, phone string) (string, error)
	ValidateOTP(c *gin.Context, req *models.ValidateOTPRequest) error
}

type authService struct {
	userRepository    users.UserRepository
	roleRepository    roles.RoleRepository
	profileRepository profile.ProfileRepository
	otpRepository     otp.OTPRepository
	WAservice         external.Whatsapp
	jwtService        JWTService
	config            *config.Configuration
	db                *gorm.DB
}

func NewAuthService(config *config.Configuration, db *gorm.DB) AuthService {
	return &authService{
		userRepository:    users.NewUserRepository(db),
		roleRepository:    roles.NewRoleRepository(db),
		profileRepository: profile.NewProfileRepository(db),
		WAservice:         *external.NewWhatsappService(config),
		otpRepository:     otp.NewOTPRepository(db),
		jwtService:        NewJWTService(config),
		db:                db,
		config:            config,
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

func (as *authService) SendOTP(c *gin.Context, phone string) (string, error) {
	tx := as.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	ctxUser := utils.GetContext(c)
	_, err := as.userRepository.GetUserByID(c, ctxUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", customError.ErrNotFound
		}
		return "", err
	}

	if err := as.checkExistingOTP(c, tx, phone); err != nil {
		tx.Rollback()
		return "", err
	}

	otp := utils.GenerateOTP()
	otpEntity := &entity.OTP{}
	otpEntity.ToEntity(phone, otp, ctxUser.ID, utils.OTPExpired())

	_, err = as.otpRepository.CreateOTP(c, tx, otpEntity)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	_, err = as.WAservice.SendMessage(phone, otp)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return otp, nil

}

func (as *authService) checkExistingOTP(c *gin.Context, tx *gorm.DB, phone string) error {

	otp, err := as.otpRepository.GetOTP(c, phone, utils.GetContext(c).ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if otp == nil {
		return nil
	}

	return as.otpRepository.DeleteOTP(c, tx, otp.ID)
}

func (as *authService) ValidateOTP(c *gin.Context, req *models.ValidateOTPRequest) error {
	ctxUser := utils.GetContext(c)
	tx := as.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	otp, err := as.otpRepository.GetOTP(c, req.Phone, ctxUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}
		return err
	}

	if err := as.validateOTP(c, tx, *ctxUser, req); err != nil {
		tx.Rollback()
		return err
	}

	if !utils.IsValidOTP(otp.OTP, req.UserOTP) || utils.IsExpiredOTP(otp.ExpiredAt) {
		tx.Rollback()
		return customError.ErrInvalidOTP
	}

	if err := as.otpRepository.DeleteOTP(c, tx, otp.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

func (as *authService) validateOTP(c *gin.Context, tx *gorm.DB, ctxUser utils.Context, req *models.ValidateOTPRequest) error {

	if req.Phone != "" {
		profile, err := as.profileRepository.GetProfileByPhone(c, req.Phone)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}

		if profile.UserID != ctxUser.ID {
			return customError.ErrNotFound
		}

		if err := as.profileRepository.UpdateVerifiedPhone(c, tx, ctxUser.ID, true); err != nil {
			return err
		}
	}

	return nil

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
