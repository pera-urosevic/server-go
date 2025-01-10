package controller

import (
	"net/http"
	"server/api/dj/query"

	"github.com/gin-gonic/gin"
)

func Query(r *gin.Engine) {
	r.GET("/dj/query/:query", func(c *gin.Context) {
		q := c.Param("query")
		res := query.Query(q)
		c.JSON(http.StatusOK, res)
	})
}
