package cmd

import (
	"os/exec"
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
		return "", err
	}
	result := strings.Trim(string(cmd_output), "\r\n")
	return result, err
}
