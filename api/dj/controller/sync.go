package controller

import (
	"net/http"
	"server/api/dj/sync"

	"github.com/gin-gonic/gin"
)

func Sync(r *gin.Engine) {
	r.PUT("/dj/sync", func(c *gin.Context) {
		res := sync.Sync()
		c.JSON(http.StatusOK, res)
	})
}
