package handler

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	config *config.Config
}

// NewCartHandler takes gin, service and config parameters and returns a new handler struct.
func NewHealthCheckHandler(r *gin.RouterGroup, cfg *config.Config) {
	hch := &HealthCheckHandler{
		config: cfg,
	}
	// Endpoints with relative functions and methods

	// GET
	r.GET("/", hch.ShowInfo)

}

func (hch *HealthCheckHandler) ShowInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      hch.config.ServerConfig.Status,
		"environment": hch.config.ServerConfig.Mode,
		"version":     hch.config.ServerConfig.AppVersion,
	})
}
