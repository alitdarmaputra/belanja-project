package user

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Login(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}
