package controller

import (
	"net/http"
	"server/api/dj/search"

	"github.com/gin-gonic/gin"
)

func Search(r *gin.Engine) {
	r.GET("/dj/search/:search", func(c *gin.Context) {
		s := c.Param("search")
		res := search.Search(s)
		c.JSON(http.StatusOK, res)
	})
}
