package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/router"
	"github.com/alitdarmaputra/belanja-project/config"
	"github.com/alitdarmaputra/belanja-project/config/db"
	orderRepository "github.com/alitdarmaputra/belanja-project/modules/database/order"
	productRepository "github.com/alitdarmaputra/belanja-project/modules/database/product"
	userRepository "github.com/alitdarmaputra/belanja-project/modules/database/user"
	shipperService "github.com/alitdarmaputra/belanja-project/modules/shipper"

	orderService "github.com/alitdarmaputra/belanja-project/bussiness/order"
	productService "github.com/alitdarmaputra/belanja-project/bussiness/product"
	userService "github.com/alitdarmaputra/belanja-project/bussiness/user"

	orderController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/order"
	productController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/product"
	userController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"

	"github.com/gin-gonic/gin"
)

const (
	production = "production"
)

func InitializeServer() *http.Server {
	cfg := config.LoadConfigAPI("./config")

	db, err := db.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if cfg.Env == production {
		gin.SetMode(gin.ReleaseMode)
	}

	userRepository := userRepository.NewUserRepository()
	productRepository := productRepository.NewProductRepository()
	orderRepository := orderRepository.NewOrderRepository()

	middleware := middleware.NewAuthentication(cfg.JWTSecretKey)

	userService := userService.NewUserService(userRepository, db)
	userService.SetJWTConfig(
		cfg.JWTSecretKey,
		time.Duration(cfg.JWTExpiredTime)*time.Minute,
	)

	userController := userController.NewUserController(userService, middleware)

	productService := productService.NewProductService(productRepository, userRepository, db)
	productController := productController.NewProductController(productService, middleware)

	shipperService := shipperService.NewShipperService(cfg.Shipper.BaseUrl, cfg.Shipper.Key)
	orderService := orderService.NewOrderService(
		shipperService,
		userRepository,
		productRepository,
		orderRepository,
		db,
	)
	orderController := orderController.NewOrderController(orderService, middleware)
	handler := router.NewRouter(userController, productController, orderController, middleware, cfg)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}

	return &server
}
