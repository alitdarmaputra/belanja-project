package router

import (
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/order"
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/product"
	"github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/config"
	"github.com/alitdarmaputra/belanja-project/constant"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController user.UserController,
	productController product.ProductController,
	orderController order.OrderController,
	authentication middleware.Authetication,
	cfg *config.Api,
) *gin.Engine {
	r := gin.New()

	r.Use(gin.CustomRecovery(middleware.ErrorHandler))

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.POST("/auth/login", userController.Login)
	v1.POST("/auth/register", userController.Create)

	v1JWTAuth := v1.Use(middleware.JWTMiddlewareAuth(cfg.JWTSecretKey))

	v1JWTAuth.PUT("/profile",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionUpdateUser,
		),
		userController.Update)

	v1JWTAuth.GET("/profile",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionShowUser,
		),
		userController.FindById)

	v1JWTAuth.POST("/products",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionCreateProduct),
		productController.Create)

	v1JWTAuth.PUT("/products",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionUpdateProduct,
			constant.PermissionShowProduct,
		),
		productController.Update)

	v1JWTAuth.DELETE("/products/:id",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionDeleteProduct,
			constant.PermissionShowProduct,
		),
		productController.Delete)

	v1JWTAuth.GET("/products/:id",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionShowProduct,
		),
		productController.FindById)

	v1JWTAuth.GET("/products",
		middleware.PermissionMiddleware(
			authentication,
			constant.PermissionShowProduct,
		),
		productController.FindAll)

	v1JWTAuth.DELETE("/order/:id", orderController.CancelOrder)
	v1JWTAuth.POST("/order", orderController.CreateOrder)
	return r
}
