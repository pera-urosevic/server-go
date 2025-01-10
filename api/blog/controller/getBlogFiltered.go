package controller

import (
	"net/http"
	"server/api/blog/database"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetBlogFiltered(r *gin.Engine) {
	r.GET("/blog/:filter", func(c *gin.Context) {
		filter := c.Param("filter")

		posts, err := database.GetPosts(filter)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, posts)
	})
}
