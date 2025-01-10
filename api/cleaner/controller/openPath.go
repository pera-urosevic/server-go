package controller

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"server/system"
	"strings"

	"github.com/gin-gonic/gin"
)

func OpenPath(r *gin.Engine) {
	r.GET("/cleaner/open/*path", func(c *gin.Context) {
		path := c.Param("path")[1:]
		root := os.Getenv("CLEANER_ROOT")
		path = root + path
		path = strings.ReplaceAll(path, "/", "\\")

		cmd := os.Getenv("CLEANER_OPEN")
		if cmd == "" {
			err := fmt.Errorf("CLEANER_OPEN not set")
			system.GinError(c, err, true)
			return
		}

		app := exec.Command(cmd, path)
		err := app.Start()
		if err == nil {
			app.Process.Release()
			c.JSON(http.StatusOK, nil)
			return
		}

		system.GinError(c, err, true)
	})
}
