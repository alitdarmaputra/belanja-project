package user

import (
	"context"

	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/cmd/api/response"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request request.UserUpdateRequest, userId int) response.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
	Login(ctx context.Context, request request.UserLoginRequest) *Token
}
