package transactions

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(c context.Context, tx *gorm.DB, transaction *entity.Transaction) (
		*entity.Transaction, error)
}

type tranasctionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &tranasctionRepository{db: db}
}

func (tr *tranasctionRepository) CreateTransaction(c context.Context, tx *gorm.DB,
	transaction *entity.Transaction) (*entity.Transaction, error) {

	if err := tx.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
