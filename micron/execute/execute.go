package execute

import (
	"os/exec"
	"server/micron/log"
)

func WithOutput(cmd string, args []string) (string, error) {
	run := exec.Command(cmd, args...)
	output, err := run.CombinedOutput()
	if err != nil {
		log.Log(err)
	}
	return string(output), err
}

func WithDetach(cmd string, args []string) error {
	app := exec.Command(cmd, args...)
	err := app.Start()
	if err == nil {
		app.Process.Release()
	}
	log.Log(err)
	return err
}
