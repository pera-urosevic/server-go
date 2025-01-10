package charged

import (
	"fmt"
	"path"
	"server/charged/cmd"
	"server/charged/log"
	"server/charged/state"
	"server/env"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func alert(name string, value int) {
	message := fmt.Sprintf("Device %s: %d%%", name, value)
	log.Log("[ALERT]", message)

	icon := path.Join(env.ExeDir(), "assets", "charged.png")
	err := beeep.Notify("Charged Alert", message, icon)
	if err != nil {
		log.Log(err)
	}
}

func tick() {
	state := state.Get()

	for i, device := range state.Devices {
		res, err := cmd.Run(device.Command)
		if err != nil {
			log.Log(err)
			continue
		}

		value, err := strconv.Atoi(res)
		if err != nil {
			log.Log(err)
			continue
		}

		if value <= device.Low && device.Value > device.Low {
			alert(device.Name, value)
		}

		state.Devices[i].Value = value
	}
}

func Start() {
	state.Load()
	tick()

	ticker := time.NewTicker(time.Duration(state.Time()) * time.Second)
	defer ticker.Stop()

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				tick()
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
