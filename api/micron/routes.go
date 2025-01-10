package micron

import (
	"server/api/micron/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.Status(r)
	controller.Pause(r)
	controller.Resume(r)
	controller.GetConfig(r)
	controller.UpdateConfig(r)
}
