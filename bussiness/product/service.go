package product

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
)

type ProductService interface {
	Create(
		ctx context.Context,
		request request.ProductCreateRequest,
		userId int,
	) response.ProductResponse
	Update(
		ctx context.Context,
		request request.ProductUpdateRequest,
		userId int,
	) response.ProductResponse
	Delete(ctx context.Context, productId, userId int)
	FindById(ctx context.Context, productId int) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
