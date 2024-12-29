package micron

import (
	"io"
	"net/http"
	"os"
	"server/api/micron/log"
	"server/api/micron/runner"
	"server/api/micron/storage"
	"server/api/micron/types"
	"server/system"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/micron/status", func(c *gin.Context) {
		status := types.Status{Enabled: active}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/pause", func(c *gin.Context) {
		Pause()
		status := types.Status{Enabled: active}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/resume", func(c *gin.Context) {
		Resume()
		status := types.Status{Enabled: active}
		c.JSON(http.StatusOK, status)
	})

	r.GET("/micron/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, storage.Config)
	})

	r.PUT("/micron/config", func(c *gin.Context) {
		payload, err := io.ReadAll(c.Request.Body)
		if err != nil {
			system.GinError(c, err, false)
		}
		storage.ConfigParse(payload)
		storage.Config.Changed = true
		storage.ConfigSave()
		c.JSON(http.StatusOK, storage.Config)
	})

	return r
}

var active = true

func Pause() {
	log.Log("[PAUSE]")
	active = false
}

func Resume() {
	log.Log("[RESUME]")
	active = true
}

func tick() {
	runner.Monitor(&storage.Config)
	runner.Daily(&storage.Config)
	runner.Weekly(&storage.Config)
	storage.ConfigSave()
}

func Start() {
	intervalString := os.Getenv("MICRON_INTERVAL")
	log.Log("[START]", intervalString+"s")

	storage.ConfigLoad()
	tick()

	interval, err := strconv.Atoi(intervalString)
	if err != nil {
		interval = 1
	}
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if active {
					tick()
				}
			case <-quit:
				log.Log("[STOP]")
				ticker.Stop()
				return
			}
		}
	}()

	for range ticker.C {
	}
}
