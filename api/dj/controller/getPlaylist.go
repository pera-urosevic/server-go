package controller

import (
	"net/http"
	"server/api/dj/playlist"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetPlaylist(r *gin.Engine) {
	r.GET("/dj/playlist/:query", func(c *gin.Context) {
		q := c.Param("query")

		res, err := playlist.Playlist(q)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
