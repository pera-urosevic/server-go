package controller

import (
	"net/http"
	"server/api/dj/search"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Search(r *gin.Engine) {
	r.GET("/dj/search/:search", func(c *gin.Context) {
		s := c.Param("search")

		res, err := search.Search(s)
		if err != nil {
			system.GinError(c, nil, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
