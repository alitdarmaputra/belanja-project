package shipper

import "context"

type ShipperService interface {
	CreateOrder(ctx context.Context, request ShipperCreateRequest) string
	CancelOrder(ctx context.Context, shipperId string)
}
