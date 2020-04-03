package tools

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

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

func ExecShellAdmin(s string, outputProcessor func(string, *os.ProcessState)) {
	script := fmt.Sprintf("osascript -e \"do shell script \\\"%s\\\" with administrator privileges\"", s)
	ExecShell(script, outputProcessor)
}

func ExecShell(s string, outputProcessor func(string, *os.ProcessState)) {
	cmd := exec.Command("/bin/bash", "-c", s + " 2>&1" )
	out, _ := cmd.StdoutPipe()
	ch := processOut(out)
	cmd.Start()
	for str := range ch {
		outputProcessor(str, nil)
	}
	cmd.Wait()
	outputProcessor("", cmd.ProcessState)
}
