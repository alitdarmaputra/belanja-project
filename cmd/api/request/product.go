package request

type ProductCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Qty  int    `json:"qty"  binding:"required,numeric,gte=0"`
}

type ProductUpdateRequest struct {
	Id   int    `json:"id"   binding:"required,numeric"`
	Name string `json:"name" binding:"required"`
	Qty  int    `json:"qty"  binding:"required,numeric,gte=0"`
}
