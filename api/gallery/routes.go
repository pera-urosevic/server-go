package gallery

import (
	"server/api/gallery/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetGallery(r)
	controller.Sync(r)
	controller.Deploy(r)
	controller.OpenPhoto(r)
	controller.GetPhoto(r)
	controller.UpdatePhoto(r)
}
