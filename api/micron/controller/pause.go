package controller

import (
	"net/http"
	"server/micron"
	"server/micron/types"

	"github.com/gin-gonic/gin"
)

func Pause(r *gin.Engine) {
	r.GET("/micron/pause", func(c *gin.Context) {
		micron.Pause()
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})
}
