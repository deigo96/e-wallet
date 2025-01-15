package entity

import (
	"time"

	"github.com/deigo96/e-wallet.git/app/models"
)

type Profile struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	FullName        string `json:"full_name"`
	Address         string `json:"address"`
	PhoneNumber     string `json:"phone"`
	PlaceOfBirth    string `json:"place_of_birth"`
	DateOfBirth     string `json:"date_of_birth"`
	UserID          int    `json:"user_id"`
	IsVerifiedPhone bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (p *Profile) ToModel() models.ProfileResponse {
	return models.ProfileResponse{
		FullName:        p.FullName,
		Address:         p.Address,
		PhoneNumber:     p.PhoneNumber,
		PlaceOfBirth:    p.PlaceOfBirth,
		DateOfBirth:     p.DateOfBirth,
		UserID:          p.UserID,
		IsVerifiedPhone: p.IsVerifiedPhone,
	}
}

func (p *Profile) ToEntity(profile models.ProfileRequest) {

	p.FullName = profile.FullName
	p.Address = profile.Address
	p.PhoneNumber = profile.PhoneNumber
	p.PlaceOfBirth = profile.PlaceOfBirth
	p.DateOfBirth = profile.DateOfBirth
	p.UserID = profile.UserID
}
