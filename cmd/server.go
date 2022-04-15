package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/handler"
	repo "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/repository"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/service"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
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

	// Router groups
	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	productRouter := rootRouter.Group("/")
	categoryRouter := rootRouter.Group("/category")
	userRouter := rootRouter.Group("/user")

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

	// Run localhost:8080
	r.Run(":8080")

}
