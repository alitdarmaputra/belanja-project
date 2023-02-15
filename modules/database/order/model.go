package order

import (
	"time"

	"github.com/alitdarmaputra/belanja-project/modules/database/product"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"gorm.io/gorm"
)

type Order struct {
	Id        int
	UsersId   int
	User      user.User `gorm:"foreignkey:UsersId"`
	OutletsId int
	OUtlet    user.User `gorm:"foreignkey:OutletsId"`
	ShipperId string
	Status    string
	Products  []product.Product `gorm:"many2many:order_details;joinForeignKey:OrdersId;joinReferences:ProductsId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
