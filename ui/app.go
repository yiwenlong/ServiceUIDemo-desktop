package ui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/yiwenlong/launchduidemo/controller"
	"os"
)

type LaunchdUIApp struct {
	app        *widgets.QApplication
	mainWindow *MainWindow
	systemTray *SystemTray
	serverCtl  controller.IServerController
}

func NewApp() *LaunchdUIApp {
	app := LaunchdUIApp{
		app: widgets.NewQApplication(len(os.Args), os.Args),
	}
	app.mainWindow = NewMainWindow(&app)
	app.systemTray = NewSystemTray(&app)
	app.serverCtl = controller.NewServerController(app.app.ApplicationDirPath())
	return &app
}

func (lapp *LaunchdUIApp) Launch() {
	lapp.mainWindow.Launch()
	lapp.app.Exec()
}

func (lapp *LaunchdUIApp) Exit() {
	lapp.app.Exit(0)
}

func (lapp *LaunchdUIApp) ShowMainWindow() {
	if lapp.mainWindow == nil {
		return
	}
	lapp.mainWindow.Show()
}

func (lapp *LaunchdUIApp) LaunchSystemTray() {
	lapp.systemTray.Launch()
}

func (lapp *LaunchdUIApp) CloseSystemTray() {
	lapp.systemTray.Close()
}
