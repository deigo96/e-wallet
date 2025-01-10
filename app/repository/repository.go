package repository

import (
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"gorm.io/gorm"
)

type Repository struct {
	userRepository users.UserRepository
	db             *gorm.DB
}

// type Repository interface {
// 	// users.UserRepository
// }

func NewRepository(
	db *gorm.DB,
) *Repository {
	return &Repository{
		db:             db,
		userRepository: users.NewUserRepository(db),
	}
}
