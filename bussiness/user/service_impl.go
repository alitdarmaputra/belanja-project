package user

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"github.com/alitdarmaputra/belanja-project/utils"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	utils.PanicIfError(err)

	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.Save(ctx, tx, user.User{
		Email:       request.Email,
		FullName:    request.FullName,
		Password:    string(hash),
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

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	utils.PanicIfError(err)

	user.Email = request.Email
	user.FullName = request.FullName
	user.PhoneNumber = request.PhoneNumber
	user.Address = request.Address
	user.AreaId = request.AreaId
	user.RoleId = request.RoleId
	user.Latitude = request.Latitude
	user.Longitude = request.Longitude

	user, err = service.UserRepository.Update(ctx, tx, user)
	utils.PanicIfError(err)

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	utils.PanicIfError(err)

	err = service.UserRepository.Delete(ctx, tx, user.Id)
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
