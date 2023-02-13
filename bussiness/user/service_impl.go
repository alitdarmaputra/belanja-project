package user

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"github.com/alitdarmaputra/belanja-project/utils"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository user.UserRepository
	DB             *gorm.DB
}

func NewUserService(
	userRepository user.UserRepository,
	db *gorm.DB,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Create(
	ctx context.Context,
	request request.UserCreateRequest,
) response.UserResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.Save(ctx, tx, user.User{
		Email:       request.Email,
		FullName:    request.FullName,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		AreaId:      request.AreaId,
		RoleId:      request.RoleId,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
	})
	utils.PanicIfError(err)

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(
	ctx context.Context,
	request request.UserUpdateRequest,
) response.UserResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.Update(ctx, tx, user.User{
		Id:          request.Id,
		Email:       request.Email,
		FullName:    request.FullName,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		AreaId:      request.AreaId,
		RoleId:      request.RoleId,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
	})
	utils.PanicIfError(err)

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	err := service.UserRepository.Delete(ctx, tx, userId)
	utils.PanicIfError(err)
}

func (service *UserServiceImpl) FindById(
	ctx context.Context,
	userId int,
) response.UserResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	utils.PanicIfError(err)

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return response.ToUserResponses(users)
}
