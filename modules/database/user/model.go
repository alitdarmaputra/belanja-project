package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int
	Email       string
	FullName    string
	Password    string
	PhoneNumber string
	Address     string
	AreaId      int
	RoleId      int8
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
