package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/database"
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

}
