package runner

import (
	"time"

	"server/api/micron/check"
	"server/api/micron/execute"
	"server/api/micron/log"
	"server/api/micron/types"
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
		} else {
			log.Log("[WEEKLY] [SUCCESS]", task.Name)
			config.Weekly[i].LastRun = now
			config.Changed = true
		}
	}
}
