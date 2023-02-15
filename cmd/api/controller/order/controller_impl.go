package order

import (
	"context"
	"net/http"

	"github.com/alitdarmaputra/belanja-project/bussiness/order"
	"github.com/alitdarmaputra/belanja-project/cmd/api/common/response"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderService order.OrderService
	Middleware   middleware.Authetication
}

func NewOrderController(
	orderService order.OrderService,
	middleware middleware.Authetication,
) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
		Middleware:   middleware,
	}
}

func (controller *OrderControllerImpl) CreateOrder(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	orderCreateRequest := request.OrderCreateRequest{}
	err = ctx.ShouldBindJSON(&orderCreateRequest)
	utils.PanicIfError(err)

	orderResponse := controller.OrderService.Create(context.TODO(), orderCreateRequest, claims.Id)
	response.JsonBasicData(ctx, http.StatusCreated, "Created", orderResponse)
}

func (controller *OrderControllerImpl) CancelOrder(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	orderId := request.PathParam{}
	err = ctx.ShouldBindUri(&orderId)
	utils.PanicIfError(err)

	controller.OrderService.Delete(ctx, orderId.Id, claims.Id)
	response.JsonBasicResponse(ctx, http.StatusOK, "OK")
}
