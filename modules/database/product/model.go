package product

import (
	"time"

	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"gorm.io/gorm"
)

type Product struct {
	Id        int            `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Qty       int            `gorm:"column:qty"`
	Price     int            `gorm:"column:price"`
	OutletsId int            `gorm:"column:outlets_id"`
	Outlet    user.User      `gorm:"foreignkey:OutletsId"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
