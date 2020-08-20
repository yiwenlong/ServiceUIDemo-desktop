package helper

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"io"
)

func RunCommand(cmd string, outputProcessor func(string)) bool {
	command := BuildCommand(cmd)
	if outputProcessor != nil {
		out, err := command.StdoutPipe()
		if err != nil {
			outputProcessor(err.Error())
		}
		ch := outChan(out)
		if err := command.Start(); err != nil {
			outputProcessor(err.Error())
			return false
		}
		for echo := range ch {
			outputProcessor(echo)
		}
	} else {
		if err := command.Start(); err != nil {
			return false
		}
	}
	if err := command.Wait(); err != nil {
		if outputProcessor != nil {
			outputProcessor(err.Error())
		}
		return false
	}
	return command.ProcessState.Success()
}

func outChan(reader io.ReadCloser) chan string {
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

func OpenLogFile(url string) {
	qurl := core.NewQUrl3(fmt.Sprintf("file://%s", url), 0)
	desktopService := gui.NewQDesktopServicesFromPointer(nil)
	desktopService.OpenUrl(qurl)
}
