package runner

import (
	"time"

	"server/micron/check"
	"server/micron/execute"
	"server/micron/log"
	"server/micron/types"
)

func Daily(config *types.Config) {
	for i, task := range config.Daily {

		if !task.Enabled {
			continue
		}

		if !check.DayTime(task.Time, task.LastRun) {
			continue
		}

		if task.Net {
			if !check.Net() {
				continue
			}
		}

		now := time.Now().Format(time.RFC3339)
		err := execute.WithDetach(task.Cmd, task.Args)
		if err != nil {
			log.Log("[DAILY] [FAIL]", task.Name, err)
			return
		}

		log.Log("[DAILY] [SUCCESS]", task.Name)
		config.Daily[i].LastRun = now
		config.Changed = true
	}
}
