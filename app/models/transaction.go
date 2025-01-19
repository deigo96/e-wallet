package models

import (
	"github.com/deigo96/e-wallet.git/app/constant"
)

type TransactionRequest struct {
	Amount          float64                  `json:"amount" validate:"required,gte=0"`
	TransactionType constant.TransactionType `json:"-"`
	UserID          int                      `json:"-"`
	Note            string                   `json:"note"`
}
