package controller

import (
	"net/http"
	"os"
	"server/api/cleaner/database"
	"server/api/cleaner/filesystem"
	"server/system"

	"github.com/gin-gonic/gin"
)

func GetUnknown(r *gin.Engine) {
	r.GET("/cleaner/unknown", func(c *gin.Context) {
		paths, err := database.GetPaths()
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		root := os.Getenv("CLEANER_ROOT")

		unknown, err := filesystem.Unknown(paths, root)
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, unknown)
	})
}
