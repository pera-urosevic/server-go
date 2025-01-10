package controller

import (
	"net/http"
	"server/api/cleaner/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func RemoveBookmark(r *gin.Engine) {
	r.DELETE("/cleaner/bookmarks", func(c *gin.Context) {
		var bookmark = ""
		err := c.BindJSON(&bookmark)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.RemoveBookmark(bookmark)
		if err != nil {
			system.GinError(c, err, true)
		}

		bookmarks, err := database.GetBookmarks()
		if err != nil {
			system.GinError(c, err, true)
		}

		c.JSON(http.StatusOK, bookmarks)
	})
}
