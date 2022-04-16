package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/cmd/server"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/logger"
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

	server.InitServer(cfg)
}
