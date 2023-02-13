package user

import (
	"context"
	"net/http"

	"github.com/alitdarmaputra/belanja-project/bussiness/user"
	"github.com/alitdarmaputra/belanja-project/cmd/api/common/response"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService user.UserService
	Middleware  middleware.Authetication
}

func NewUserController(
	userService user.UserService,
	middleware middleware.Authetication,
) UserController {
	return &UserControllerImpl{
		UserService: userService,
		Middleware:  middleware,
	}
}

func (controller *UserControllerImpl) Create(ctx *gin.Context) {
	userCreateRequest := request.UserCreateRequest{}
	err := ctx.ShouldBindJSON(&userCreateRequest)
	utils.PanicIfError(err)

	userResponse := controller.UserService.Create(context.Background(), userCreateRequest)
	response.JsonBasicData(ctx, http.StatusCreated, "Created", userResponse)
}

func (controller *UserControllerImpl) Update(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	userUpdateRequest := request.UserUpdateRequest{}
	err = ctx.ShouldBindJSON(&userUpdateRequest)
	utils.PanicIfError(err)

	userResponse := controller.UserService.Update(
		context.Background(),
		userUpdateRequest,
		claims.Id,
	)
	response.JsonBasicData(ctx, http.StatusOK, "OK", userResponse)
}

func (controller *UserControllerImpl) Delete(ctx *gin.Context) {
	pathParam := request.PathParam{}

	err := ctx.ShouldBindUri(&pathParam)
	utils.PanicIfError(err)

	controller.UserService.Delete(context.Background(), pathParam.Id)
	response.JsonBasicResponse(ctx, http.StatusOK, "OK")
}

func (controller *UserControllerImpl) FindById(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	userResponse := controller.UserService.FindById(context.Background(), claims.Id)
	response.JsonBasicData(ctx, http.StatusOK, "OK", userResponse)
}

func (controller *UserControllerImpl) FindAll(ctx *gin.Context) {
	userResponses := controller.UserService.FindAll(context.Background())
	response.JsonBasicData(ctx, http.StatusOK, "OK", userResponses)
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {
	userLoginRequest := request.UserLoginRequest{}
	err := ctx.ShouldBind(&userLoginRequest)
	utils.PanicIfError(err)

	token := controller.UserService.Login(context.Background(), userLoginRequest)

	response.JsonBasicData(ctx, http.StatusOK, "OK", token)
}
