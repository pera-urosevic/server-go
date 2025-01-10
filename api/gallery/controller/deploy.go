package controller

import (
	"net/http"
	"os"
	"os/exec"
	"server/system"
	"strings"

	"github.com/gin-gonic/gin"
)

func Deploy(r *gin.Engine) {
	r.GET("/gallery/deploy", func(c *gin.Context) {
		deploy := os.Getenv("GALLERY_DEPLOY")
		cmd := strings.Split(deploy, "|")

		app := exec.Command(cmd[0], cmd[1:]...)
		err := app.Start()
		app.Process.Release()
		if err != nil {
			system.GinError(c, err, true)
			return
		}

		c.JSON(http.StatusOK, nil)
	})
}
