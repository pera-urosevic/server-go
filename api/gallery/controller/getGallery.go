package controller

import (
	"net/http"
	"server/api/gallery/database"
	"server/system"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetGallery(r *gin.Engine) {
	r.GET("/gallery/:sort/:pixelfed/*filter", func(c *gin.Context) {
		sort := c.Param("sort")
		sort = strings.ToLower(sort)
		if (sort != "asc") && (sort != "desc") {
			sort = "desc"
		}

		pixelfed := c.Param("pixelfed")
		pixelfed = strings.ToLower(pixelfed)
		pixelfedUpload := pixelfed == "true"

		filter := c.Param("filter")[1:]

		photos, err := database.GetPhotos(filter, sort, pixelfedUpload)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		c.JSON(http.StatusOK, photos)
	})
}
