package execute

import (
	"os/exec"
)

func WithOutput(cmd string, args []string) (string, error) {
	run := exec.Command(cmd, args...)
	output, err := run.CombinedOutput()
	return string(output), err
}

func WithDetach(cmd string, args []string) error {
	app := exec.Command(cmd, args...)
	err := app.Start()
	if err == nil {
		app.Process.Release()
	}
	return err
}
