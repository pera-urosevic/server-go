package controller

import (
	"net/http"
	"server/api/gallery/database"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPhoto(r *gin.Engine) {
	r.GET("/gallery/photo/:photoID", func(c *gin.Context) {
		photoID, err := strconv.ParseInt(c.Param("photoID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		post, err := database.GetPhoto(photoID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, post)
	})
}
