package controller

import (
	"net/http"
	"server/api/blog/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetBlog(r *gin.Engine) {
	r.GET("/blog/", func(c *gin.Context) {

		posts, err := database.GetPosts("")
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, posts)
	})
}
