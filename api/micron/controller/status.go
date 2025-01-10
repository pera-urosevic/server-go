package controller

import (
	"net/http"
	"server/micron"
	"server/micron/types"

	"github.com/gin-gonic/gin"
)

func Status(r *gin.Engine) {
	r.GET("/micron/status", func(c *gin.Context) {
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})
}
