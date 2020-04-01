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

type MainWindow struct {
	lapp			*LaunchdApp
	mWindow 		*widgets.QMainWindow
	loggerWidget 	*widgets.QTextEdit
	centralWidget 	*widgets.QWidget
	btnStart 		*widgets.QPushButton
	btnClose 		*widgets.QPushButton
	btnShowLog 		*widgets.QPushButton
}

type SystemTray struct {
	lapp			*LaunchdApp
	mSystemTrayIcon	*widgets.QSystemTrayIcon
}

type LaunchdApp struct {
	application 	*widgets.QApplication
	mainWindow 		*MainWindow
	systemTray		*SystemTray
}

func (la *LaunchdApp) asyncLog(reader io.ReadCloser) error {
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
			la.mainWindow.loggerWidget.Append(fmt.Sprintf("%s%s", cache, line))
			cache = s[len(s)-1]
		}
	}
	return nil
}

func (la *LaunchdApp) execShell(script string) error {
	cmd := exec.Command("/bin/bash", "-c", script)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	go la.asyncLog(stderr)
	go la.asyncLog(stdout)
	err := cmd.Wait()
	return err
}

func (la *LaunchdApp) showDialog(s string, widget *widgets.QWidget) {
	messageBox := widgets.NewQMessageBox2(widgets.QMessageBox__NoIcon, "Notification", s, widgets.QMessageBox__Ok, widget, core.Qt__WindowTitleHint)
	messageBox.Show()
}

func (la *LaunchdApp) appDir() *core.QDir {
	return core.NewQDir2(la.application.ApplicationDirPath())
}

func (la *LaunchdApp) startServer() error {
	dir := la.appDir()
	dir.Cd("../../")
	startScript := dir.AbsoluteFilePath("start.sh")
	go la.execShell(fmt.Sprintf("%s %s", startScript, dir.AbsolutePath()))
	return nil
}

func (la *LaunchdApp) stopServer() error {
	dir := la.appDir()
	dir.Cd("../../")
	stopScript := dir.AbsoluteFilePath("stop.sh")

	go la.execShell(fmt.Sprintf("%s %s", stopScript, dir.AbsolutePath()))
	return nil
}

func (la *LaunchdApp) showServerLog() {
	homeDir := la.appDir()
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

func (la *LaunchdApp) launch() {

}

func (m *MainWindow) setUp() {
	m.mWindow = widgets.NewQMainWindow(nil, 0)
	m.mWindow.SetMinimumSize2(800, 400)
	m.mWindow.SetWindowTitle("launchd demo")

	m.centralWidget = widgets.NewQWidget(nil, 0)
	m.centralWidget.SetLayout(widgets.NewQVBoxLayout())
	m.centralWidget.Layout().QLayoutItem.SetAlignment(core.Qt__AlignLeft)
	m.mWindow.SetCentralWidget(m.centralWidget)

	m.loggerWidget = widgets.NewQTextEdit(nil)
	m.loggerWidget.SetReadOnly(true)

	m.btnStart = widgets.NewQPushButton2("start launchd test server", nil)
	m.btnStart.ConnectClicked(func(bool) {
		if err := m.lapp.startServer(); err != nil {
			m.lapp.showDialog(err.Error(), m.centralWidget)
		}
	})
	m.centralWidget.Layout().AddWidget(m.btnStart)

	m.btnClose = widgets.NewQPushButton2("close launchd test server", nil)
	m.btnClose.ConnectClicked(func(bool) {
		if err := m.lapp.stopServer(); err != nil {
			m.lapp.showDialog(err.Error(), m.centralWidget)
		}
	})
	m.centralWidget.Layout().AddWidget(m.btnClose)

	m.btnShowLog = widgets.NewQPushButton2("view server log", nil)
	m.btnShowLog.ConnectClicked(func(bool) {
		m.lapp.showServerLog()
	})
	m.centralWidget.Layout().AddWidget(m.btnShowLog)

	m.centralWidget.Layout().AddWidget(m.loggerWidget)
}

func (m *MainWindow) show() {
	if m.mWindow == nil {
		m.setUp()
	}
	if m.mWindow.IsVisible() {
		return
	}
	m.mWindow.Show()
}

func (s *SystemTray) setUp() {
	s.mSystemTrayIcon = widgets.NewQSystemTrayIcon(nil)
	s.mSystemTrayIcon.SetIcon(widgets.NewQCommonStyle().StandardIcon(widgets.QStyle__SP_MessageBoxCritical, nil, nil))
	menu := widgets.NewQMenu(nil)
	exit := menu.AddAction("Exit")
	exit.ConnectTriggered(func(bool) {
		s.lapp.exit()
	})
	s.mSystemTrayIcon.SetContextMenu(menu)
}

func (s *SystemTray) show() {
	if s.mSystemTrayIcon == nil {
		s.setUp()
	}
	if s.mSystemTrayIcon.IsVisible() {
		return
	}
	s.mSystemTrayIcon.Show()
}

func (la *LaunchdApp) setUp() {
	la.application = widgets.NewQApplication(len(os.Args), os.Args)
	la.mainWindow = &MainWindow{
		lapp: la,
	}
	la.mainWindow.setUp()
	la.systemTray = &SystemTray{
		lapp: la,
	}
	la.systemTray.setUp()
}

func (la *LaunchdApp) show() {
	la.setUp()
	la.systemTray.show()
	la.mainWindow.show()
	la.application.Exec()
}

func (la *LaunchdApp) exit() {
	la.application.Exit(0)
}

func main() {
	launchdApp := LaunchdApp{}
	launchdApp.show()
}
