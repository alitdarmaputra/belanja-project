package product

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/modules/database"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(
	ctx context.Context,
	tx *gorm.DB,
	product Product,
) (Product, error) {
	result := tx.Create(&product)
	return product, database.WrapError(result.Error)
}

func (repository *ProductRepositoryImpl) Update(
	ctx context.Context,
	tx *gorm.DB,
	product Product,
) (Product, error) {
	result := tx.Save(&product)
	return product, database.WrapError(result.Error)
}

func (repository *ProductRepositoryImpl) Delete(
	ctx context.Context,
	tx *gorm.DB,
	productId int,
) error {
	result := tx.Delete(&Product{}, productId)
	return database.WrapError(result.Error)
}

func (repository *ProductRepositoryImpl) FindById(
	ctx context.Context,
	tx *gorm.DB,
	productId int,
) (Product, error) {
	var product Product
	result := tx.First(&product, productId)
	return product, database.WrapError(result.Error)
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []Product {
	var products []Product
	tx.Find(&products)
	return products
}
