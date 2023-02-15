package order

import (
	"context"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *gorm.DB, order Order) (Order, error)
	Update(ctx context.Context, tx *gorm.DB, order Order) (Order, error)
	FindById(ctx context.Context, tx *gorm.DB, orderId, userId int) (Order, error)
	FindAll(ctx context.Context, tx *gorm.DB, userId int) []Order
}
