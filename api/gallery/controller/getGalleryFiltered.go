package controller

import (
	"net/http"
	"server/api/gallery/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetGalleryFiltered(r *gin.Engine) {
	r.GET("/gallery/:filter", func(c *gin.Context) {
		filter := c.Param("filter")

		photos, err := database.GetPhotos(filter)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		c.JSON(http.StatusOK, photos)
	})
}
