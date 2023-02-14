package product

import (
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *gorm.DB, product Product) (Product, error)
	Update(ctx context.Context, tx *gorm.DB, user Product) (Product, error)
	Delete(ctx context.Context, tx *gorm.DB, userId int) error
	FindById(ctx context.Context, tx *gorm.DB, userId int) (Product, error)
	FindAll(ctx context.Context, tx *gorm.DB) []Product
}
