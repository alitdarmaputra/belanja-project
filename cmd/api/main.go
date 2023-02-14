package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	productService "github.com/alitdarmaputra/belanja-project/bussiness/product"
	userService "github.com/alitdarmaputra/belanja-project/bussiness/user"
	productController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/product"
	userController "github.com/alitdarmaputra/belanja-project/cmd/api/controller/user"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/router"
	"github.com/alitdarmaputra/belanja-project/config/db"
	productRepository "github.com/alitdarmaputra/belanja-project/modules/database/product"
	userRepository "github.com/alitdarmaputra/belanja-project/modules/database/user"
)

func main() {
	db, err := db.NewMySQL()
	if err != nil {
		panic(err)
	}
	userRepository := userRepository.NewUserRepository()
	productRepository := productRepository.NewProductRepository()

	middleware := middleware.NewAuthentication("default-secret-key")

	userService := userService.NewUserService(userRepository, db)
	userController := userController.NewUserController(userService, middleware)

	productService := productService.NewProductService(productRepository, userRepository, db)
	productController := productController.NewProductController(productService, middleware)

	handler := router.NewRouter(userController, productController)

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
