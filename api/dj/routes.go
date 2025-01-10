package dj

import (
	"server/api/dj/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.Search(r)
	controller.GetQueries(r)
	controller.Query(r)
	controller.GetPlaylist(r)
	controller.Sync(r)
}
