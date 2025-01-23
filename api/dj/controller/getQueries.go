package controller

import (
	"net/http"
	"server/api/dj/query"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetQueries(r *gin.Engine) {
	r.GET("/dj/queries", func(c *gin.Context) {
		res, err := query.Queries()
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
