package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/logger"
)

func main() {
	fmt.Println("E-Commerce service starting...")

	// Initialize zap logger.
	logger.InitLogger()
	defer logger.Close()
}
