package controller

import (
	"net/http"
	"server/api/dj/query"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Query(r *gin.Engine) {
	r.GET("/dj/query/:query", func(c *gin.Context) {
		q := c.Param("query")

		res, err := query.Query(q)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
