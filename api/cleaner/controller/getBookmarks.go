package controller

import (
	"net/http"
	"server/api/cleaner/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetBookmarks(r *gin.Engine) {
	r.GET("/cleaner/bookmarks", func(c *gin.Context) {
		bookmarks, err := database.GetBookmarks()
		if err != nil {
			system.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, bookmarks)
	})
}
