package controller

import (
	"net/http"
	"server/api/gallery/database"
	"server/api/gallery/flick"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Flickr(r *gin.Engine) {
	r.GET("/gallery/photo/:photoID/flickr", func(c *gin.Context) {
		photoID, err := strconv.ParseInt(c.Param("photoID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		photo, err := database.GetPhoto(photoID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		res, err := flick.Upload(photo)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
