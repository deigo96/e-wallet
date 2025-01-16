package profile

import (
	"errors"

	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/external"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/otp"
	"github.com/deigo96/e-wallet.git/app/repository/profile"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/services/auth"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfileService interface {
	GetProfile(c *gin.Context) (*models.ProfileResponse, error)
	CreateProfile(c *gin.Context, request *models.ProfileRequest) (*models.ProfileResponse, error)
	UpdateProfile(c *gin.Context, request *models.ProfileRequest) (*models.ProfileResponse, error)
}

type profileService struct {
	profileRepository profile.ProfileRepository
	userRepository    users.UserRepository
	otpRepository     otp.OTPRepository
	authService       auth.AuthService
	waService         external.Whatsapp
	config            *config.Configuration
	db                *gorm.DB
}

func NewProfileService(config *config.Configuration, db *gorm.DB) ProfileService {
	return &profileService{
		profileRepository: profile.NewProfileRepository(db),
		userRepository:    users.NewUserRepository(db),
		authService:       auth.NewAuthService(config, db),
		otpRepository:     otp.NewOTPRepository(db),
		waService:         *external.NewWhatsappService(config),
		config:            config,
		db:                db,
	}
}

func (ps *profileService) GetProfile(c *gin.Context) (*models.ProfileResponse, error) {
	userID := utils.GetID(c)

	profile, err := ps.profileRepository.GetProfile(c, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	user, err := ps.userRepository.GetUserByID(c, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	response := profile.ToModel()

	response.Email = user.Email
	response.IsActive = user.IsActive

	return &response, nil
}

func (ps *profileService) CreateProfile(c *gin.Context, request *models.ProfileRequest) (*models.ProfileResponse, error) {
	userID := utils.GetID(c)

	request.PhoneNumber = utils.RefactorPhoneNumber(request.PhoneNumber)

	tx := ps.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := ps.userRepository.GetUserByID(c, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	existingProfile, err := ps.GetProfile(c)
	if err != nil && !errors.Is(err, customError.ErrNotFound) {
		return nil, err
	}

	if existingProfile != nil {
		return nil, customError.ErrProfileAlreadyCreated
	}

	request.UserID = user.ID
	profileEntity := entity.Profile{}
	profileEntity.ToEntity(*request)

	profile, err := ps.profileRepository.CreateProfile(c, tx, &profileEntity)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	otp := utils.GenerateOTP()

	otpEntity := entity.OTP{}
	otpEntity.ToEntity(profile.PhoneNumber, otp, profile.UserID, utils.OTPExpired())
	_, err = ps.otpRepository.CreateOTP(c, tx, &otpEntity)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = ps.waService.SendMessage(profile.PhoneNumber, otp)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response := profile.ToModel()

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &response, nil
}

func (ps *profileService) UpdateProfile(c *gin.Context, request *models.ProfileRequest) (*models.ProfileResponse, error) {
	ctxUser := utils.GetContext(c)

	profile, err := ps.profileRepository.GetProfile(c, ctxUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	profileEntity := &entity.Profile{}
	profileEntity.ToEntity(*request)
	profileEntity.ID = profile.ID

	res, err := ps.profileRepository.UpdateProfile(c, profileEntity)
	if err != nil {
		return nil, err
	}

	profileResponse := res.ToModel()

	return &profileResponse, nil
}
