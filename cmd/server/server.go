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

func InitServer(cfg *config.Config) {

	// Database connection
	DB := database.Connect(cfg)
	defer database.Close(DB)

	// Set gin mode
	gin.SetMode(gin.ReleaseMode)

	// Gin initialize
	r := gin.Default()

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

	// Service and Handlers
	productService := service.NewProductService(repo)
	handler.NewProductHandler(productRouter, productService, cfg)

	categoryService := service.NewCategoryService(repo)
	handler.NewCategoryHandler(categoryRouter, categoryService)

	userService := service.NewUserService(repo)
	handler.NewUserHandler(userRouter, userService, cfg)

	cartService := service.NewCartService(repo)
	handler.NewCartHandler(cartRouter, cartService, cfg)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs))

	// Run localhost:8080
	// r.Run(":8080")

}
