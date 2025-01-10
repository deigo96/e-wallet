package users

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(c context.Context) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetAllUsers(c context.Context) ([]entity.User, error) {
	users := []entity.User{}

	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
