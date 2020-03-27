package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func asyncLog(reader io.ReadCloser, logger *widgets.QTextEdit) error {
	cache := ""
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n")
			logger.Append(fmt.Sprintf("%s%s", cache, line))
			cache = s[len(s)-1]
		}
	}
	return nil
}

func execShell(logger *widgets.QTextEdit, script string) error {
	cmd := exec.Command("/bin/bash", "-c", script)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	go asyncLog(stderr, logger)
	go asyncLog(stdout, logger)
	err := cmd.Wait()
	return err
}

func show(s string, widget *widgets.QWidget) {
	messageBox := widgets.NewQMessageBox2(widgets.QMessageBox__NoIcon, "Notification", s, widgets.QMessageBox__Ok, widget, core.Qt__WindowTitleHint)
	messageBox.Show()
}

func startServer(app *widgets.QApplication, logger *widgets.QTextEdit) error {
	appDirPath := app.ApplicationDirPath()
	dir2 := core.NewQDir2(appDirPath)
	dir2.Cd("../../")
	startScript := dir2.AbsoluteFilePath("start.sh")
	go execShell(logger, fmt.Sprintf("%s %s", startScript, dir2.AbsolutePath()))
	return nil
}

func stopServer(app *widgets.QApplication, logger *widgets.QTextEdit) error {
	appDirPath := app.ApplicationDirPath()
	dir2 := core.NewQDir2(appDirPath)
	dir2.Cd("../../")
	stopScript := dir2.AbsoluteFilePath("stop.sh")

	go execShell(logger, fmt.Sprintf("%s %s", stopScript, dir2.AbsolutePath()))
	return nil
}

func showServerLog(app *widgets.QApplication) {
	appDirPath := app.ApplicationDirPath()
	homeDir := core.NewQDir2(appDirPath)
	homeDir.Cd("../../")
	logFile := core.NewQFile2(homeDir.AbsoluteFilePath("error.log"))
	logFile.Open(core.QIODevice__ReadOnly)
	file := logFile.ReadAll()
	logFile.Close()
	textBrowser := widgets.NewQTextBrowser(nil)
	textBrowser.SetMinimumSize2(600, 600)
	textBrowser.SetWindowTitle(logFile.FileName())
	textBrowser.SetText(file.Data())
	textBrowser.Show()
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 400)
	window.SetWindowTitle("launchd demo")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.Layout().QLayoutItem.SetAlignment(core.Qt__AlignLeft)
	window.SetCentralWidget(widget)

	logTextWidget := widgets.NewQTextEdit(nil)
	logTextWidget.SetReadOnly(true)

	buttonStart := widgets.NewQPushButton2("启动 launchd 测试服务", nil)
	buttonStart.ConnectClicked(func(bool) {
		if err := startServer(app, logTextWidget); err != nil {
			show(err.Error(), widget)
		}
	})
	widget.Layout().AddWidget(buttonStart)

	buttonClose := widgets.NewQPushButton2("close launchd test server", nil)
	buttonClose.ConnectClicked(func(bool) {
		if err := stopServer(app, logTextWidget); err != nil {
			show(err.Error(), widget)
		}
	})
	widget.Layout().AddWidget(buttonClose)

	buttonShowLog := widgets.NewQPushButton2("view server log", nil)
	buttonShowLog.ConnectClicked(func(bool) {
		showServerLog(app)
	})
	widget.Layout().AddWidget(buttonShowLog)
	widget.Layout().AddWidget(logTextWidget)
	window.Show()
	app.Exec()
}
