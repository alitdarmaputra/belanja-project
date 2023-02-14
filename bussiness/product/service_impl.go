package product

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
	"github.com/alitdarmaputra/belanja-project/modules/database/product"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"github.com/alitdarmaputra/belanja-project/utils"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository product.ProductRepository
	UserRepository    user.UserRepository
	DB                *gorm.DB
}

func NewProductService(
	productRepository product.ProductRepository,
	userRepository user.UserRepository,
	db *gorm.DB,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		DB:                db,
	}
}

func (service *ProductServiceImpl) Create(
	ctx context.Context,
	request request.ProductCreateRequest,
	userId int,
) response.ProductResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	outlet, err := service.UserRepository.FindById(ctx, tx, userId)
	utils.PanicIfError(err)

	product, err := service.ProductRepository.Save(ctx, tx, product.Product{
		Name:   request.Name,
		Qty:    request.Qty,
		Outlet: outlet,
	})
	utils.PanicIfError(err)

	return response.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(
	ctx context.Context,
	request request.ProductUpdateRequest,
	userId int,
) response.ProductResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	utils.PanicIfError(err)

	if product.OutletsId != userId {
		utils.PanicIfError(bussiness.NewUnauthorizedError("Unauthorized"))
	}

	product.Name = request.Name
	product.Qty = request.Qty

	product, err = service.ProductRepository.Update(ctx, tx, product)
	utils.PanicIfError(err)

	return response.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId, userId int) {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	utils.PanicIfError(err)

	if product.OutletsId != userId {
		utils.PanicIfError(bussiness.NewUnauthorizedError("Unauthorized"))
	}

	err = service.ProductRepository.Delete(ctx, tx, product.Id)
	utils.PanicIfError(err)
}

func (service *ProductServiceImpl) FindById(
	ctx context.Context,
	productId int,
) response.ProductResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	utils.PanicIfError(err)

	return response.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return response.ToProductResponses(products)
}
