package helper

import (
	"os/exec"
	"syscall"
)

func BuildCommand(cmd string) *exec.Cmd {
	c := exec.Command("cmd", "/c", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return c
}
