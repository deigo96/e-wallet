package entity

import "github.com/deigo96/e-wallet.git/app/models"

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ToModel() models.User {
	return models.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
