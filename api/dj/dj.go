package dj

import (
	"net/http"
	"server/api/dj/playlist"
	"server/api/dj/query"
	"server/api/dj/search"
	"server/api/dj/sync"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/dj/search/:search", func(c *gin.Context) {
		s := c.Param("search")
		res := search.Search(s)
		c.JSON(http.StatusOK, res)
	})

	r.GET("/dj/queries", func(c *gin.Context) {
		res := query.Queries()
		c.JSON(http.StatusOK, res)
	})

	r.GET("/dj/query/:query", func(c *gin.Context) {
		q := c.Param("query")
		res := query.Query(q)
		c.JSON(http.StatusOK, res)
	})

	r.GET("/dj/playlist/:query", func(c *gin.Context) {
		q := c.Param("query")
		res := playlist.Playlist(q)
		c.JSON(http.StatusOK, res)
	})

	r.PUT("/dj/sync", func(c *gin.Context) {
		res := sync.Sync()
		c.JSON(http.StatusOK, res)
	})

	return r
}
