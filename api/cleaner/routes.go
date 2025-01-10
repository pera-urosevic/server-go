package cleaner

import (
	"server/api/cleaner/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetBookmarks(r)
	controller.AddBookmark(r)
	controller.RemoveBookmark(r)
	controller.OpenPath(r)
	controller.GetPath(r)
	controller.AddPath(r)
	controller.RemovePath(r)
	controller.DeleteEntry(r)
}
