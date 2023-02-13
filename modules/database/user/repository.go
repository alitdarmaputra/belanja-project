package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user User) (User, error)
	Update(ctx context.Context, tx *gorm.DB, user User) (User, error)
	Delete(ctx context.Context, tx *gorm.DB, userId int) error
	FindById(ctx context.Context, tx *gorm.DB, userId int) (User, error)
	FindAll(ctx context.Context, tx *gorm.DB) []User
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (User, error)
}
