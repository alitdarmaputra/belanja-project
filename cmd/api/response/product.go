package response

import (
	"github.com/alitdarmaputra/belanja-project/modules/database/product"
)

type ProductResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func ToProductResponse(product product.Product) ProductResponse {
	return ProductResponse{
		Id:   product.Id,
		Name: product.Name,
		Qty:  product.Qty,
	}
}

func ToProductResponses(products []product.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
