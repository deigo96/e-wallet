package users

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(c context.Context) ([]entity.User, error)
	CreateUser(c context.Context, user *entity.User) error
	GetUserFilter(c context.Context, username, email, operator string) (*entity.User, error)
	GetUserByID(c context.Context, userID int) (*entity.User, error)
	GetUserByUsername(c context.Context, username string) (*entity.User, error)
	GetUserByEmail(c context.Context, email string) (*entity.User, error)
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

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*entity.User, error) {
	user := &entity.User{}

	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) GetUserByUsername(c context.Context, username string) (*entity.User, error) {
	user := &entity.User{}

	if err := ur.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) GetUserByID(c context.Context, userID int) (*entity.User, error) {
	user := &entity.User{}

	if err := ur.db.First(user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) GetUserFilter(c context.Context, username, email, operator string) (*entity.User, error) {
	user := &entity.User{}

	query := ur.db

	if username != "" && email == "" {
		query = query.Where("username = ?", username)
	}

	if email != "" && username == "" {
		query = query.Where("email = ?", email)
	}

	if username != "" && email != "" {
		query = query.Where("username = ? "+operator+" email = ?", username, email)
	}

	if err := query.Debug().Table(user.TableName()).Find(user).Limit(1).Error; err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (ur *userRepository) CreateUser(c context.Context, user *entity.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return constant.ErrInternalServerError
	}

	return nil
}
