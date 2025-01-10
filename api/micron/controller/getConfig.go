package controller

import (
	"net/http"
	"server/micron/storage"

	"github.com/gin-gonic/gin"
)

func GetConfig(r *gin.Engine) {
	r.GET("/micron/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, storage.Config)
	})
}
