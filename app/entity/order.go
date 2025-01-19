package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	OrderID   string
	Item      string
	Qty       int
	Amount    decimal.Decimal
	Note      string
	CreatedAt time.Time
}
