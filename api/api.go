package api

import (
	"os"
	"server/api/blog"
	"server/api/dj"
	"server/api/gallery"
	"server/api/micron"
	"server/system"

	"github.com/gin-gonic/gin"
)

func Start() {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "release"
	}
	gin.SetMode(mode)
	gin.ForceConsoleColor()

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "55555"
	}
	addr := ":" + port

	system.Log("API", addr)
	router := system.GinCustom()
	router = micron.Routes(router)
	router = dj.Routes(router)
	router = blog.Routes(router)
	router = gallery.Routes(router)
	router.Run(addr)
}
