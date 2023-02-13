package router

import (
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/exception"
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController user.UserController) *gin.Engine {
	r := gin.New()

	r.Use(gin.CustomRecovery(exception.ErrorHandler))

	v1 := r.Group("/v1")
	v1.PUT("/users", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)
	v1.GET("/users", userController.FindAll)
	v1.GET("/users/:id", userController.FindById)
	v1.POST("/auth/login", userController.Login)
	v1.POST("/auth/register", userController.Create)

	return r
}
