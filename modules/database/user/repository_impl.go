package user

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/modules/database"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(
	ctx context.Context,
	tx *gorm.DB,
	user User,
) (User, error) {
	result := tx.Create(&user)
	return user, database.WrapError(result.Error)
}

func (repository *UserRepositoryImpl) Update(
	ctx context.Context,
	tx *gorm.DB,
	user User,
) (User, error) {
	result := tx.Save(&user)
	return user, database.WrapError(result.Error)
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, userId int) error {
	result := tx.Delete(&User{}, userId)
	return database.WrapError(result.Error)
}

func (repository *UserRepositoryImpl) FindById(
	ctx context.Context,
	tx *gorm.DB,
	userId int,
) (User, error) {
	var user User
	result := tx.First(&user, userId)
	return user, database.WrapError(result.Error)
}

func (repository *UserRepositoryImpl) FindByEmail(
	ctx context.Context,
	tx *gorm.DB,
	email string,
) (User, error) {
	var user User
	result := tx.First(&user, "email = ?", email)
	return user, database.WrapError(result.Error)
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []User {
	var users []User
	tx.Find(&users)
	return users
}
