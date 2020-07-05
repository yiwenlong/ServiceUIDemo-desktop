package helper

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type SessionToken int

type ProcessCallback interface {
	Echo(token SessionToken, echo string)

	OnError(token SessionToken, exitCode int, state string)

	OnSuccess(token SessionToken)
}

func processOut(reader io.ReadCloser) chan string {
	out := make(chan string)
	go func() {
		buf := make([]byte, 1024)
		for {
			bCount, err := reader.Read(buf)
			if err != nil {
				break
			}
			if bCount > 0 {
				out <- string(buf[0:bCount])
			}
		}
		close(out)
	}()
	return out
}

func ExecShellAdmin(s string, handler ProcessCallback, token SessionToken) {
	script := fmt.Sprintf("osascript -e \"do shell script \\\"%s\\\" with administrator privileges\"", s)
	ExecShellAsync(script, handler, token)
}

func ExecShellAsync(s string, handler ProcessCallback, token SessionToken) {
	cmd := exec.Command("/bin/bash", "-c", s+" 2>&1")
	out, _ := cmd.StdoutPipe()
	ch := processOut(out)
	cmd.Start()
	for echo := range ch {
		handler.Echo(token, echo)
	}
	cmd.Wait()
	state := cmd.ProcessState
	if state.Success() {
		handler.OnSuccess(token)
	} else {
		handler.OnError(token, state.ExitCode(), state.String())
	}
}

func ExecShell(s string) (*os.ProcessState, string) {
	cmd := exec.Command("/bin/bash", "-c", s+" 2>&1")
	out, _ := cmd.Output()
	cmd.Run()
	return cmd.ProcessState, string(out)
}
