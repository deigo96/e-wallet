package profile

import (
	"context"
	"time"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(c context.Context, userID int) (*entity.Profile, error)
	CreateProfile(c context.Context, tx *gorm.DB, profile *entity.Profile) (*entity.Profile, error)
	GetProfileByPhone(c context.Context, phone string) (*entity.Profile, error)
	UpdateVerifiedPhone(c context.Context, tx *gorm.DB, userID int, isVerifiedPhone bool) error
	UpdateProfile(c context.Context, profile *entity.Profile) (*entity.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (pr *profileRepository) GetProfile(c context.Context, userID int) (*entity.Profile, error) {
	profile := &entity.Profile{}

	if err := pr.db.Where("user_id = ?", userID).First(profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func (pr *profileRepository) CreateProfile(c context.Context, tx *gorm.DB, profile *entity.Profile) (*entity.Profile, error) {

	err := tx.Create(&profile).Error
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (pr *profileRepository) GetProfileByPhone(c context.Context, phone string) (*entity.Profile, error) {

	profile := &entity.Profile{}

	if err := pr.db.Where("phone_number = ?", phone).First(profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func (pr *profileRepository) UpdateProfile(c context.Context, profile *entity.Profile) (*entity.Profile, error) {

	err := pr.db.Model(&entity.Profile{}).
		Where("id = ?", profile.ID).
		Updates(map[string]interface{}{
			"full_name":      profile.FullName,
			"place_of_birth": profile.PlaceOfBirth,
			"date_of_birth":  profile.DateOfBirth,
			"address":        profile.Address,
			"updated_at":     time.Now(),
		}).Error

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (pr *profileRepository) UpdateVerifiedPhone(c context.Context, tx *gorm.DB, userID int, isVerifiedPhone bool) error {

	return tx.Model(&entity.Profile{}).Where("user_id = ?", userID).Update("is_verified_phone", isVerifiedPhone).Error

}
