package users

import (
	"errors"
	"log"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/external"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/balances"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers(c *gin.Context) ([]models.User, error)
	CreateUser(c *gin.Context, user *models.CreateUserRequest) error
	VerifyEmail(c *gin.Context, email, token string) error
	ResendEmailVerification(c *gin.Context) error
}

type userService struct {
	userRepository    users.UserRepository
	emailService      external.EmailService
	balanceRepository balances.BalanceRepository
	config            *config.Configuration
	db                *gorm.DB
}

func NewUserService(config *config.Configuration, db *gorm.DB) UserService {
	return &userService{
		userRepository:    users.NewUserRepository(db),
		emailService:      external.NewEmailService(config),
		balanceRepository: balances.NewBalanceRepository(db),
		config:            config,
		db:                db}
}

func (us userService) GetAllUsers(c *gin.Context) ([]models.User, error) {
	ctxValue := utils.GetContext(c)

	if !utils.IsAdmin(ctxValue.Role) {
		return nil, customError.ErrUnauthorized
	}

	users, err := us.userRepository.GetAllUsers(c)
	if err != nil {
		return nil, err
	}

	userResponse := []models.User{}
	for _, user := range users {
		userResponse = append(userResponse, user.ToModel(constant.GetRoleValue(ctxValue.Role)))
	}

	return userResponse, nil
}

func (us *userService) CreateUser(c *gin.Context, user *models.CreateUserRequest) error {
	tx := us.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userEmailResponse, err := us.userRepository.GetUserByEmail(c, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if userEmailResponse != nil {
		if userEmailResponse.EmailVerification != nil && !userEmailResponse.IsActive {
			return us.resendEmailVerification(c, userEmailResponse.ID)
		}
		return customError.ErrEmailAlreadyUsed
	}

	userNameResponse, err := us.userRepository.GetUserByUsername(c, user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if userNameResponse != nil {
		return customError.ErrUsernameAlreadyUsed
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	emailVerification := utils.GenerateEmailVerification(user.Email)

	userEntity := entity.User{
		Username:          user.Username,
		Email:             user.Email,
		Password:          password,
		Role:              constant.ROLE_USER,
		EmailVerification: &emailVerification,
	}

	ctx := utils.GetContext(c)

	if ctx.ID != 0 && (ctx.Role == constant.GetRoleName(constant.ROLE_ADMIN) ||
		ctx.Role == constant.GetRoleName(constant.ROLE_SUPER_ADMIN)) {

		userEntity.CreatedBy = ctx.ID
		userEntity.UpdatedBy = ctx.ID
	}

	if err := us.userRepository.CreateUser(c, tx, &userEntity); err != nil {
		tx.Rollback()
		return err
	}

	if err := us.sendEmailVerification(&userEntity); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

func (us *userService) VerifyEmail(c *gin.Context, email, token string) error {
	tx := us.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := us.userRepository.GetUserByEmail(c, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}
		return err
	}

	if user.EmailVerification == nil {
		return customError.ErrNotFound
	}

	if *user.EmailVerification != token {
		return customError.ErrNotFound
	}

	balance := &entity.Balance{
		UserID:  user.ID,
		Balance: 0,
	}

	_, err = us.balanceRepository.CreateBalance(c, tx, balance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := us.userRepository.ActivateUser(c, email); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (us *userService) ResendEmailVerification(c *gin.Context) error {
	ctxUser := utils.GetContext(c)

	return us.resendEmailVerification(c, ctxUser.ID)
}

func (us *userService) resendEmailVerification(c *gin.Context, userID int) error {

	user, err := us.userRepository.GetUserByID(c, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}
		return err
	}

	if user.EmailVerification == nil {
		log.Println("Email verification token not found")
		return customError.ErrNotFound
	}

	if user.IsActive {
		log.Println("User is already active")
		return customError.ErrNotFound
	}

	if err := us.sendEmailVerification(user); err != nil {
		return err
	}

	return nil
}

func (us *userService) sendEmailVerification(user *entity.User) error {
	if user.Email == "" && user.EmailVerification == nil {
		return customError.ErrNotFound
	}

	linkEmail := "https://" + utils.GenerateLinkEmailVerification(us.config, user.Email)

	message := `
		<!DOCTYPE html>
		<html>
		<body>
		<p>Hi [Recipient Name],</p>
		<p>Thank you for joining us. To get started, please click the link below:</p>
		<p><a href="` + linkEmail + `">Click here to verify your email</a></p>
		<p>If you have any questions, feel free to reach out to us at support@example.com.</p>
		<p>Best regards,<br>Your Team</p>
		</body>
		</html>
		`

	if err := us.emailService.SendEmail(
		user.Email,
		"Email verification",
		message); err != nil {
		return err
	}

	return nil
}
