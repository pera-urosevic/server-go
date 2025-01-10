package api

import (
	"os"
	"server/api/blog"
	"server/api/cleaner"
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

	host := os.Getenv("API_HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "55555"
	}
	addr := host + ":" + port

	system.Log("API", addr)

	router := system.GinCustom()

	blog.Routes(router)
	cleaner.Routes(router)
	dj.Routes(router)
	gallery.Routes(router)
	micron.Routes(router)

	router.Run(addr)
}
