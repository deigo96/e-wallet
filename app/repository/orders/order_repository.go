package orders

import (
	"context"

	"github.com/deigo96/e-wallet.git/app/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(c context.Context, tx *gorm.DB, order *entity.Order) (*entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (or *orderRepository) CreateOrder(c context.Context, tx *gorm.DB, order *entity.Order) (
	*entity.Order, error) {

	if err := tx.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
