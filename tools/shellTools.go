package tools

import (
	"io"
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

func ExecShell(s string, outputProcessor func(string, bool)) {
	cmd := exec.Command("/bin/bash", "-c", s + " 2>&1" )
	out, _ := cmd.StdoutPipe()
	ch := processOut(out)
	cmd.Start()
	for str := range ch {
		outputProcessor(str, false)
	}
	outputProcessor("", true)
}