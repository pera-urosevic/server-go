package controller

import (
	"net/http"
	"server/api/gallery/sync"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Sync(r *gin.Engine) {
	r.GET("/gallery/sync", func(c *gin.Context) {
		photos, err := sync.Sync()
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		c.JSON(http.StatusOK, photos)
	})
}
