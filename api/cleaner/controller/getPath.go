package controller

import (
	"server/api/cleaner/controller/lib"

	"github.com/gin-gonic/gin"
)

func GetPath(r *gin.Engine) {
	r.GET("/cleaner/list/*path", func(c *gin.Context) {
		path := c.Param("path")[1:]
		lib.List(c, path)
	})
}
