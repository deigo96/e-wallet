package transactions

import (
	"context"
	"time"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(c context.Context, tx *gorm.DB, transaction *entity.Transaction) (
		*entity.Transaction, error)
	GetLastTransaction(c context.Context, userID int) (*entity.Transaction, error)
	UpdateTransactionStatus(c context.Context, tx *gorm.DB, transaction *entity.Transaction) (
		*entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (tr *transactionRepository) CreateTransaction(c context.Context, tx *gorm.DB,
	transaction *entity.Transaction) (*entity.Transaction, error) {

	if err := tx.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (tr *transactionRepository) GetLastTransaction(c context.Context, userID int) (
	*entity.Transaction, error) {

	var transaction entity.Transaction
	err := tr.db.Where("user_id = ?", userID).Last(&transaction).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &transaction, nil
}

func (tr *transactionRepository) UpdateTransactionStatus(c context.Context, tx *gorm.DB,
	transaction *entity.Transaction) (*entity.Transaction, error) {

	query := map[string]interface{}{
		"status":     transaction.Status,
		"updated_at": time.Now(),
	}

	if transaction.Status == constant.TransactionSuccess && transaction.PaidAt != nil {
		query["paid_at"] = transaction.PaidAt
	}

	if err := tx.Model(&entity.Transaction{}).
		Where("id = ?", transaction.ID).
		Updates(query).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
