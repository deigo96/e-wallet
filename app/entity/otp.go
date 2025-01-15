package entity

import "time"

type OTP struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Phone     string    `json:"phone"`
	OTP       string    `json:"otp" gorm:"otp"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (o *OTP) ToEntity(phone, otp string, userID int, expiredAt time.Time) {
	o.Phone = phone
	o.OTP = otp
	o.UserID = userID
	o.ExpiredAt = expiredAt
}
