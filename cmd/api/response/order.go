package response

import "github.com/alitdarmaputra/belanja-project/modules/database/order"

type OrderResponse struct {
	Id       int               `json:"id"`
	User     UserResponse      `json:"user"`
	Outlet   UserResponse      `json:"outlet"`
	Products []ProductResponse `json:"product"`
	Status   string            `json:"status"`
}

func ToOrderResponse(order order.Order) OrderResponse {
	return OrderResponse{
		Id:       order.Id,
		User:     ToUserResponse(order.User),
		Outlet:   ToUserResponse(order.OUtlet),
		Products: ToProductResponses(order.Products),
		Status:   order.Status,
	}
}
