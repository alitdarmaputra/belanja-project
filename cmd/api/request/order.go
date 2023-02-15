package request

import "github.com/alitdarmaputra/belanja-project/modules/shipper"

type OrderCreateRequest struct {
	Destination shipper.Destination `json:"destination" binding:"required"`
	ProductId   int                 `json:"product_id"  binding:"required"`
	Qty         int                 `json:"qty"         binding:"required"`
}
