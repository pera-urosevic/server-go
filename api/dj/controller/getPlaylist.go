package controller

import (
	"net/http"
	"server/api/dj/playlist"

	"github.com/gin-gonic/gin"
)

func GetPlaylist(r *gin.Engine) {
	r.GET("/dj/playlist/:query", func(c *gin.Context) {
		q := c.Param("query")
		res := playlist.Playlist(q)
		c.JSON(http.StatusOK, res)
	})
}
