package request

type ProductCreateRequest struct {
	Name  string `json:"name"  binding:"required"`
	Qty   int    `json:"qty"   binding:"required,numeric,gte=0"`
	Price int    `json:"price" binding:"required,numeric"`
}

type ProductUpdateRequest struct {
	Id    int    `json:"id"    binding:"required,numeric"`
	Name  string `json:"name"  binding:"required"`
	Qty   int    `json:"qty"   binding:"required,numeric,gte=0"`
	Price int    `json:"price" binding:"required,numeric"`
}
