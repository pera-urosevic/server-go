package check

import (
	"regexp"
	"strings"

	"server/micron/execute"
)

func Running(app string, re string) bool {
	appName := strings.Replace(app, ".exe", "", 1)
	cmd := "pwsh.exe"
	args := []string{"-c", "get-process '" + appName + "' | % commandline"}

	output, err := execute.WithOutput(cmd, args)
	if err != nil {
		return false
	}
	lines := strings.Split(strings.ToLower(output), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Get-Process: Cannot find a process") {
			continue
		}

		found := regexp.MustCompile("(?i)" + re).MatchString(line)
		if found {
			return true
		}
	}

	return false
}
