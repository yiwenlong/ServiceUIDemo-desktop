package helper

import (
	"fmt"
	"os/exec"
)

func BuildCommand(cmd string) *exec.Cmd {
	cmdWrapper := fmt.Sprintf("%s 2>&1", cmd)
	return exec.Command("/bin/bash", "-c", cmdWrapper)
}
