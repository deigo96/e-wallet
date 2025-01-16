package entity

import "time"

type Balance struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	Balance   float64   `json:"balance"`
	IsBlocked bool      `json:"is_blocked"`
	UpdatedAt time.Time `json:"updated_at"`
}
