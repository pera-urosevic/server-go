package runner

import (
	"server/api/micron/check"
	"server/api/micron/execute"
	"server/api/micron/log"
	"server/api/micron/types"
)

func Monitor(config *types.Config) {
	for _, task := range config.Monitor {
		if !task.Enabled {
			continue
		}
		if check.Running(task.Match.Cmd, task.Match.Regex) {
			continue
		}
		if task.Net {
			if !check.Net() {
				continue
			}
		}
		err := execute.WithDetach(task.Cmd, task.Args)
		if err != nil {
			log.Log("[MONITOR] [FAIL]", task.Name, err)
		} else {
			log.Log("[MONITOR] [SUCCESS]", task.Name)
		}
	}
}
