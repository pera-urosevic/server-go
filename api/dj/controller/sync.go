package controller

import (
	"net/http"
	"server/api/dj/sync"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Sync(r *gin.Engine) {
	r.PUT("/dj/sync", func(c *gin.Context) {
		res, err := sync.Sync()
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
