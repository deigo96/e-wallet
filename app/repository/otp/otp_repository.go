package otp

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type OTPRepository interface {
	GetOTP(c context.Context, phone string, userID int) (*entity.OTP, error)
	CreateOTP(c context.Context, tx *gorm.DB, otp *entity.OTP) (*entity.OTP, error)
	DeleteOTP(c context.Context, tx *gorm.DB, id int) error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) GetOTP(c context.Context, phone string, userID int) (*entity.OTP, error) {

	otp := &entity.OTP{}
	query := r.db.Where("user_id = ?", userID)
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	if err := query.Last(otp).Error; err != nil {
		return nil, err
	}

	return otp, nil
}

func (r *otpRepository) CreateOTP(c context.Context, tx *gorm.DB, otp *entity.OTP) (*entity.OTP, error) {

	err := tx.Create(&otp).Error
	if err != nil {
		return nil, err
	}

	return otp, nil
}

func (r *otpRepository) DeleteOTP(c context.Context, tx *gorm.DB, id int) error {

	return tx.Delete(&entity.OTP{}, id).Error
}
