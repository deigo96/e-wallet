package users

import (
	"context"
	"fmt"

	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/config"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers(c context.Context) ([]models.User, error)
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

func (us userService) GetAllUsers(c context.Context) ([]models.User, error) {
	users, err := us.userRepository.GetAllUsers(c)
	if err != nil {
		return nil, err
	}

	userResponse := []models.User{}
	for _, user := range users {
		userResponse = append(userResponse, user.ToModel())
	}
	fmt.Println(len(userResponse))
	fmt.Println(userResponse)

	return userResponse, nil
}
