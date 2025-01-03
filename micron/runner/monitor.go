package runner

import (
	"server/micron/check"
	"server/micron/execute"
	"server/micron/log"
	"server/micron/types"
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
