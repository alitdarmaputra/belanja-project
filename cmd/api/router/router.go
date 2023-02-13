package router

import (
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController user.UserController) *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1")
	v1.POST("/users", userController.Create)
	v1.PUT("/users", userController.Update)
	v1.DELETE("/users", userController.Delete)
	v1.GET("/users", userController.FindAll)
	v1.GET("/users/:id", userController.FindById)
	return r
}
