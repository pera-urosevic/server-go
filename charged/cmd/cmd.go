package cmd

import (
	"os/exec"
	"server/charged/log"
	"strings"
	"syscall"
)

func Run(cmdline string) (string, error) {
	parts := strings.Split(cmdline, " ")
	exe, args := parts[0], parts[1:]

	cmd_instance := exec.Command(exe, args...)
	cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd_output, err := cmd_instance.Output()
	if err != nil {
		log.Log(err)
		return "", err
	}

	result := strings.Trim(string(cmd_output), "\r\n")
	log.Log(err)
	return result, err
}
