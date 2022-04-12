package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/logger"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("E-Commerce service starting...")

	// Initialize zap logger.
	logger.InitLogger()
	defer logger.Close()

	cfg, err := config.LoadConfig("./../pkg/config/config-local")
	if err != nil {
		zap.L().Error("Path not found for LoadConfig()", zap.Error(err))
	}
	fmt.Println(cfg)

	DB := database.Connect(cfg)
	defer database.Close(DB)
}
