package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/handler"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/graceful"
	"github.com/gin-gonic/gin"
)

// InitServer is the function which will take configuration parameters and starts
// server connections with database.
func InitServer(cfg *config.Config) {

	// Database connection
	DB := database.Connect(cfg)
	defer database.Close(DB)

	// Set gin mode as release mode
	gin.SetMode(gin.ReleaseMode)

	// Gin initialize
	r := gin.Default()

	// Server parameters
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	// Router groups
	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	productRouter := rootRouter.Group("/")
	categoryRouter := rootRouter.Group("/category")
	userRouter := rootRouter.Group("/user")
	cartRouter := rootRouter.Group("/")

	// Repository
	repo := repo.NewRepository(DB)
	repo.Migrations()

	// Product Service and Handler
	productService := service.NewProductService(repo)
	handler.NewProductHandler(productRouter, productService, cfg)

	// Category Service and Handler
	categoryService := service.NewCategoryService(repo)
	handler.NewCategoryHandler(categoryRouter, categoryService, cfg)

	// User Service and Handler
	userService := service.NewUserService(repo)
	handler.NewUserHandler(userRouter, userService, cfg)

	// Cart Service and Handler
	cartService := service.NewCartService(repo)
	handler.NewCartHandler(cartRouter, cartService, cfg)

	// Listen and Serve
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shotdown for running processes
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs))

}
