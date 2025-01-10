package controller

import (
	"net/http"
	"server/api/cleaner/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func AddBookmark(r *gin.Engine) {
	r.POST("/cleaner/bookmarks", func(c *gin.Context) {
		var bookmark = ""
		err := c.BindJSON(&bookmark)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		err = database.AddBookmark(bookmark)
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
