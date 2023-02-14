package product

import (
	"time"

	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"gorm.io/gorm"
)

type Product struct {
	Id        int
	Name      string
	Qty       int
	Price     int
	OutletsId int
	Outlet    user.User `gorm:"foreignkey:OutletsId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
