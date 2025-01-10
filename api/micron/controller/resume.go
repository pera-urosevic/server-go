package controller

import (
	"net/http"
	"server/micron"
	"server/micron/types"

	"github.com/gin-gonic/gin"
)

func Resume(r *gin.Engine) {
	r.GET("/micron/resume", func(c *gin.Context) {
		micron.Resume()
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})
}
