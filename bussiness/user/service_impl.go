package user

import (
	"context"
	"errors"
	"time"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
	"github.com/alitdarmaputra/belanja-project/modules/database/user"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	defaultSecretKey  = "default-secret-key"
	defaultJWTExpired = 24 * time.Hour
)

type UserServiceImpl struct {
	UserRepository user.UserRepository
	DB             *gorm.DB
	jwtSecretKey   string
	jwtExpired     time.Duration
}

func NewUserService(
	userRepository user.UserRepository,
	db *gorm.DB,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		jwtSecretKey:   defaultSecretKey,
		jwtExpired:     defaultJWTExpired,
	}
}

func (service *UserServiceImpl) SetJWTConfig(secret string, expired time.Duration) {
	service.jwtSecretKey = secret
	service.jwtExpired = expired
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
	userId int,
) response.UserResponse {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
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

func (service *UserServiceImpl) Login(
	ctx context.Context,
	request request.UserLoginRequest,
) *Token {
	tx := service.DB.Begin()
	defer utils.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(bussiness.NewUnauthorizedError("Incorrect email and password entered"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		panic(bussiness.NewUnauthorizedError("Incorrect email and password entered"))
	}

	token, err := service.GenerateToken(user)
	utils.PanicIfError(err)

	return &Token{
		Token: token,
	}
}

func (service *UserServiceImpl) GenerateToken(user user.User) (string, error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   user.Id,
			"exp":  time.Now().Add(service.jwtExpired).Unix(),
			"role": user.RoleId,
		},
	)

	return eJWT.SignedString([]byte(service.jwtSecretKey))
}
