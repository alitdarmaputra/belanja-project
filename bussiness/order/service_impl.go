package order

import (
	"context"
	"fmt"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
	"github.com/alitdarmaputra/belanja-project/constant"
	"github.com/alitdarmaputra/belanja-project/modules/database/order"
	"github.com/alitdarmaputra/belanja-project/modules/database/product"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"github.com/alitdarmaputra/belanja-project/modules/shipper"
	"github.com/alitdarmaputra/belanja-project/utils"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	Shipper           shipper.ShipperService
	UserRepository    user.UserRepository
	ProductRepository product.ProductRepository
	OrderRepository   order.OrderRepository
	DB                *gorm.DB
}

func NewOrderService(
	shipper shipper.ShipperService,
	userRepository user.UserRepository,
	productRepository product.ProductRepository,
	orderRepository order.OrderRepository,
	db *gorm.DB,
) OrderService {
	return &OrderServiceImpl{
		Shipper:           shipper,
		UserRepository:    userRepository,
		ProductRepository: productRepository,
		OrderRepository:   orderRepository,
		DB:                db,
	}
}

func (service *OrderServiceImpl) Create(
	ctx context.Context,
	request request.OrderCreateRequest,
	userId int,
) response.OrderResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	item, err := service.ProductRepository.FindById(ctx, tx, request.ProductId)
	utils.PanicIfError(err)

	if item.Qty < request.Qty {
		panic(bussiness.NewBadRequestError("Qty greater than available qty"))
	}

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	utils.PanicIfError(err)

	shipperRequest := shipper.ShipperCreateRequest{
		Consignee: shipper.Consignee{
			Name:        user.FullName,
			PhoneNumber: user.PhoneNumber,
		},
		Coverage: "domestic",
		Destination: shipper.Destination{
			Address: request.Destination.Address,
			AreaId:  request.Destination.AreaId,
			Lat:     request.Destination.Lat,
			Lng:     request.Destination.Lng,
		},
		Origin: shipper.Origin{
			Address: item.Outlet.Address,
			AreaId:  item.Outlet.AreaId,
			Lat:     fmt.Sprintf("%f", item.Outlet.Latitude),
			Lng:     fmt.Sprintf("%f", item.Outlet.Longitude),
		},
		Package: shipper.Package{
			Height: 1,
			Length: 1,
			Width:  1,
			Weight: 1,
			Price:  item.Price * request.Qty,
			Items: []shipper.Item{
				{
					Name:  item.Name,
					Price: item.Price,
					Qty:   request.Qty,
				},
			},
			PackageType: 2,
		},
		Consigner: shipper.Consignee{
			Name:        item.Outlet.FullName,
			PhoneNumber: item.Outlet.PhoneNumber,
		},
		PaymentType: "postpay",
	}

	shipperId := service.Shipper.CreateOrder(ctx, shipperRequest)

	order := order.Order{}
	order.User = user
	order.OUtlet = item.Outlet
	order.ShipperId = shipperId
	order.Products = []product.Product{item}
	order.Status = constant.StatusProcess

	order, err = service.OrderRepository.Save(context.Background(), tx, order)
	utils.PanicIfError(err)

	return response.ToOrderResponse(order)
}
