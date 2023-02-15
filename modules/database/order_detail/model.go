package orderdetail

import (
	"time"

	"github.com/alitdarmaputra/belanja-project/modules/database/order"
	"github.com/alitdarmaputra/belanja-project/modules/database/product"
	"gorm.io/gorm"
)

type OrderDetail struct {
	OrdersId  int
	Order     order.Order `gorm:"foreignkey:OrdersId"`
	ProductId int
	Product   product.Product `gorm:"foreignkey:ProductsId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
