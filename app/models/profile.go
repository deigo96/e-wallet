package models

type ProfileResponse struct {
	FullName        string `json:"full_name"`
	Address         string `json:"address"`
	PhoneNumber     string `json:"phone"`
	PlaceOfBirth    string `json:"place_of_birth"`
	DateOfBirth     string `json:"date_of_birth"`
	UserID          int    `json:"user_id"`
	IsVerifiedPhone bool   `json:"is_verified_phone"`
}

type ProfileRequest struct {
	FullName     string `json:"full_name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	PhoneNumber  string `json:"phone" validate:"required"`
	PlaceOfBirth string `json:"place_of_birth" validate:"required"`
	DateOfBirth  string `json:"date_of_birth" validate:"required"`
	UserID       int    `json:"-"`
}
