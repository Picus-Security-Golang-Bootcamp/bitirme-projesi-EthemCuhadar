package middleware

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMetaData, err := helper.ExtractMetaData(c.Request, cfg)
		if err != nil {
			return
		}
		if jwtMetaData.IsAdmin {
			c.Next()
			c.Abort()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		}
	}
}

func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c.Request, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func UserMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		jwtMetaData, err := helper.ExtractMetaData(c.Request, cfg)
		if err != nil {
			return
		}
		if jwtMetaData.UserID == user_id {
			c.Next()
			c.Abort()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		}
	}
}
