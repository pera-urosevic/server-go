package micron

import (
	"os"
	"server/micron/log"
	"server/micron/runner"
	"server/micron/storage"
	"strconv"
	"time"
)

var active = true

func Pause() {
	log.Log("[PAUSE]")
	active = false
}

func Resume() {
	log.Log("[RESUME]")
	active = true
}

func IsActive() bool {
	return active
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
