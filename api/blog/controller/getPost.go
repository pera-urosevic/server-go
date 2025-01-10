package controller

import (
	"net/http"
	"server/api/blog/database"
	"server/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPost(r *gin.Engine) {
	r.GET("/blog/post/:postID", func(c *gin.Context) {
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
		c.JSON(http.StatusOK, post)
	})
}
