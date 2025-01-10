package runner

import (
	"time"

	"server/micron/check"
	"server/micron/execute"
	"server/micron/log"
	"server/micron/types"
)

func Weekly(config *types.Config) {
	for i, task := range config.Weekly {

		if !task.Enabled {
			continue
		}

		if !check.WeekDayTime(task.Day, task.Time, task.LastRun) {
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
			log.Log("[WEEKLY] [FAIL]", task.Name, err)
			return
		}

		log.Log("[WEEKLY] [SUCCESS]", task.Name)
		config.Weekly[i].LastRun = now
		config.Changed = true
	}
}
