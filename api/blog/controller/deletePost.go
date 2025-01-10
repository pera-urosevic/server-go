package controller

import (
	"net/http"
	"server/api/blog/database"
	"server/api/blog/scraper/image"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePost(r *gin.Engine) {
	r.DELETE("/blog/post/:postID", func(c *gin.Context) {
		postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
		if err != nil {
			system.GinError(c, err, false)
			return
		}

		post, err := database.GetPost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		err = image.Delete(post)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		err = database.RemovePost(postID)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
