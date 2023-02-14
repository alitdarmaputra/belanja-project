package router

import (
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/product"
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController user.UserController,
	productController product.ProductController,
) *gin.Engine {
	r := gin.New()

	r.Use(gin.CustomRecovery(middleware.ErrorHandler))

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.POST("/auth/login", userController.Login)
	v1.POST("/auth/register", userController.Create)

	v1JWTAuth := v1.Use(middleware.JWTMiddlewareAuth("default-secret-key"))
	v1JWTAuth.PUT("/profile", userController.Update)
	v1JWTAuth.GET("/profile", userController.FindById)

	v1JWTAuth.POST("/products", productController.Create)
	v1JWTAuth.PUT("/products", productController.Update)
	v1JWTAuth.DELETE("/products/:id", productController.Delete)
	v1JWTAuth.GET("/products/:id", productController.FindById)
	v1JWTAuth.GET("/products", productController.FindAll)
	return r
}
