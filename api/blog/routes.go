package blog

import (
	"server/api/blog/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetBlog(r)
	controller.GetBlogFiltered(r)
	controller.DeployBlog(r)
	controller.CreatePost(r)
	controller.GetPost(r)
	controller.UpdatePost(r)
	controller.ScrapePost(r)
	controller.DeletePost(r)
}
