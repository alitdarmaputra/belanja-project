package order

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/modules/database"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(
	ctx context.Context,
	tx *gorm.DB,
	order Order,
) (Order, error) {
	result := tx.Create(&order)
	return order, database.WrapError(result.Error)
}

func (repository *OrderRepositoryImpl) Update(
	ctx context.Context,
	tx *gorm.DB,
	order Order,
) (Order, error) {
	result := tx.Updates(&order)
	return order, database.WrapError(result.Error)
}

func (repository *OrderRepositoryImpl) FindById(
	ctx context.Context,
	tx *gorm.DB,
	orderId,
	userId int,
) (Order, error) {
	var order Order
	result := tx.Preload("User").
		Where("(outlets_id = ? OR users_id = ?) AND id = ?", userId, userId, orderId).
		First(&order)
	return order, database.WrapError(result.Error)
}

func (repository *OrderRepositoryImpl) FindAll(
	ctx context.Context,
	tx *gorm.DB,
	userId int,
) []Order {
	var orders []Order
	tx.Joins("User").Where("outlets_id = ? OR users_id = ?", userId, userId).Find(&orders)
	return orders
}
