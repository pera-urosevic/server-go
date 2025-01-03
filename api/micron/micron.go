package micron

import (
	"io"
	"net/http"
	"server/micron"
	"server/micron/storage"
	"server/micron/types"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/micron/status", func(c *gin.Context) {
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/pause", func(c *gin.Context) {
		micron.Pause()
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/resume", func(c *gin.Context) {
		micron.Resume()
		status := types.Status{Enabled: micron.IsActive()}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, storage.Config)
	})

	r.PUT("/micron/config", func(c *gin.Context) {
		payload, err := io.ReadAll(c.Request.Body)
		if err != nil {
			system.GinError(c, err, false)
		}
		storage.ConfigParse(payload)
		storage.Config.Changed = true
		storage.ConfigSave()
		c.JSON(http.StatusOK, storage.Config)
	})

	return r
}
