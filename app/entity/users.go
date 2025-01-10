package entity

import (
	"time"

	"github.com/deigo96/e-wallet.git/app/models"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	IsActive  bool
	Role      int
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ToEntity(user models.User) *User {
	u.ID = user.ID
	u.Username = user.Username
	u.Email = user.Email
	u.Password = user.Password

	return u
}

func (u *User) ToModel() models.User {
	user := models.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}

	// if u.Role == 1 || u.Role == 2 {
	// 	user.IsActive = u.IsActive
	// 	user.CreatedAt = u.CreatedAt
	// 	user.UpdatedAt = u.UpdateAt
	// }

	return user
}
