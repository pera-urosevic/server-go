package controller

import (
	"net/http"
	"server/api/gallery/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetGallery(r *gin.Engine) {
	r.GET("/gallery/", func(c *gin.Context) {
		photos, err := database.GetPhotos("")
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, photos)
	})
}
