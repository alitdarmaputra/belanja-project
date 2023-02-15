package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	orderService "github.com/alitdarmaputra/belanja-project/bussiness/order"
	productService "github.com/alitdarmaputra/belanja-project/bussiness/product"
	userService "github.com/alitdarmaputra/belanja-project/bussiness/user"
	orderController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/order"
	productController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/product"
	userController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/router"
	"github.com/alitdarmaputra/belanja-project/config"
	"github.com/alitdarmaputra/belanja-project/config/db"
	orderRepository "github.com/alitdarmaputra/belanja-project/modules/database/order"
	productRepository "github.com/alitdarmaputra/belanja-project/modules/database/product"
	userRepository "github.com/alitdarmaputra/belanja-project/modules/database/user"
	shipperService "github.com/alitdarmaputra/belanja-project/modules/shipper"
	"github.com/gin-gonic/gin"
)

const (
	production = "production"
)

func main() {
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
		Addr:    "localhost:3000",
		Handler: handler,
	}

	go func() {
		//Service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen and serve failed:", err.Error())
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Panicln("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("shutting down failed:", err.Error())
	}

	log.Println("server exiting")
}
