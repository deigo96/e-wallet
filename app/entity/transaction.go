package entity

import (
	"time"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID              int `gorm:"primaryKey"`
	UserID          int
	OrderID         string
	TotalAmount     decimal.Decimal
	TransactionType constant.TransactionType
	Note            string
	Status          constant.TransactionStatus
	CreatedAt       time.Time
	CreatedBy       *string
	UpdatedAt       time.Time
	UpdatedBy       *string
	PaidAt          *time.Time
}
