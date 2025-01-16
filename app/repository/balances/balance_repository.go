package balances

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	GetBalanceUser(c context.Context, userID int) (*entity.Balance, error)
	UpdateBalance(c context.Context, tx *gorm.DB, balance *entity.Balance) (*entity.Balance, error)
	CreateBalance(c context.Context, tx *gorm.DB, balance *entity.Balance) (*entity.Balance, error)
}

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) BalanceRepository {
	return &balanceRepository{db: db}
}

func (r *balanceRepository) GetBalanceUser(c context.Context, userID int) (*entity.Balance, error) {
	var balance entity.Balance

	if err := r.db.Where("user_id = ?", userID).First(&balance).Error; err != nil {
		return nil, err
	}

	return &balance, nil
}

func (r *balanceRepository) UpdateBalance(c context.Context, tx *gorm.DB, balance *entity.Balance) (*entity.Balance, error) {
	if err := tx.Model(&entity.Balance{}).
		Where("user_id = ?", balance.UserID).
		Update("balance", balance.Balance).Error; err != nil {
		return nil, err
	}

	return balance, nil
}

func (r *balanceRepository) CreateBalance(c context.Context, tx *gorm.DB, balance *entity.Balance) (
	*entity.Balance, error) {
	if err := tx.Create(&balance).Error; err != nil {
		return nil, err
	}

	return balance, nil
}
