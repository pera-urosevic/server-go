package controller

import (
	"net/http"
	"server/api/dj/query"

	"github.com/gin-gonic/gin"
)

func GetQueries(r *gin.Engine) {
	r.GET("/dj/queries", func(c *gin.Context) {
		res := query.Queries()
		c.JSON(http.StatusOK, res)
	})
}
