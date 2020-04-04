package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ShellToken int

type ShellHandler interface{
	HandleEcho(token ShellToken, echo string)
	HandleError(token ShellToken, exitCode int, state string)
	HandleSuccess(token ShellToken)
}

func processOut(reader io.ReadCloser) chan string{
	out := make(chan string)
	go func() {
		buf := make([]byte, 1024)
		for {
			bcount, err := reader.Read(buf)
			if err != nil {
				break
			}
			if bcount > 0 {
				out <- string(buf[0: bcount])
			}
		}
		close(out)
	}()
	return out
}

func ExecShellAdmin(s string, handler ShellHandler, token ShellToken) {
	script := fmt.Sprintf("osascript -e \"do shell script \\\"%s\\\" with administrator privileges\"", s)
	ExecShellAsync(script, handler, token)
}

func ExecShellAsync(s string, handler ShellHandler, token ShellToken) {
	cmd := exec.Command("/bin/bash", "-c", s + " 2>&1" )
	out, _ := cmd.StdoutPipe()
	ch := processOut(out)
	cmd.Start()
	for echo := range ch {
		handler.HandleEcho(token, echo)
	}
	cmd.Wait()
	state := cmd.ProcessState
	if state.Success() {
		handler.HandleSuccess(token)
	} else {
		handler.HandleError(token, state.ExitCode(), state.String())
	}
}

func ExecShell(s string) (*os.ProcessState, string)  {
	cmd := exec.Command("/bin/bash", "-c", s + " 2>&1" )
	out, _ := cmd.Output()
	cmd.Run()
	return cmd.ProcessState, string(out)
}
