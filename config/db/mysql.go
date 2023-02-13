package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL() (db *gorm.DB, err error) {
	dsn := "root:root123@tcp(localhost:3306)/belanja_project_db?charset=utf8mb4&parseTime=true&loc=Local"

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	if db, err = gorm.Open(mysql.Open(dsn), gormConfig); err != nil {
		return
	}
	return
}
