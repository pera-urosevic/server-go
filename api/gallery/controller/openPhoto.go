package controller

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func OpenPhoto(r *gin.Engine) {
	r.GET("/gallery/open/:photoPath", func(c *gin.Context) {
		photoPath := c.Param("photoPath")

		app := exec.Command("explorer.exe", photoPath)
		err := app.Start()
		app.Process.Release()

		c.JSON(http.StatusOK, err)
	})
}
