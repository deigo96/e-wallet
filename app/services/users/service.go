package users

import (
	"errors"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers(c *gin.Context) ([]models.User, error)
	CreateUser(c *gin.Context, user *models.CreateUserRequest) error
}

type userService struct {
	userRepository users.UserRepository
	config         *config.Configuration
}

func NewUserService(config *config.Configuration, db *gorm.DB) UserService {
	return &userService{
		userRepository: users.NewUserRepository(db),
		config:         config}
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

	userEmailResponse, err := us.userRepository.GetUserByEmail(c, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if userEmailResponse != nil {
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

	userEntity := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: password,
		Role:     constant.ROLE_USER,
	}

	ctx := utils.GetContext(c)

	if ctx.ID != 0 && (ctx.Role == constant.GetRoleName(constant.ROLE_ADMIN) ||
		ctx.Role == constant.GetRoleName(constant.ROLE_SUPER_ADMIN)) {

		userEntity.CreatedBy = ctx.ID
		userEntity.UpdatedBy = ctx.ID
	}

	if err := us.userRepository.CreateUser(c, &userEntity); err != nil {
		return err
	}

	return nil

}
