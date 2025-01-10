package controller

import (
	"io"
	"net/http"
	"server/micron/storage"
	"server/system"

	"github.com/gin-gonic/gin"
)

func UpdateConfig(r *gin.Engine) {
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
}
