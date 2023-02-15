package order

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
)

type OrderService interface {
	Create(
		ctx context.Context,
		request request.OrderCreateRequest,
		userId int,
	) response.OrderResponse
	Delete(
		ctx context.Context,
		orderId,
		userId int,
	)
}
