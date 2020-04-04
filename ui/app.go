package ui

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

type LaunchdUIApp struct {
	app 		*widgets.QApplication
	mainWindow 	*MainWindow
	systemTray 	*SystemTray
}

func NewApp() *LaunchdUIApp {
	app := LaunchdUIApp{
		app:        widgets.NewQApplication(len(os.Args), os.Args),
	}
	app.mainWindow = NewMainWindow(&app)
	app.systemTray = NewSystemTray(&app)
	return &app
}

func (lapp *LaunchdUIApp) Launch()  {
	lapp.systemTray.Launch()
	lapp.app.Exec()
}

func (lapp *LaunchdUIApp) Exit()  {
	lapp.app.Exit(0)
}

func (lapp *LaunchdUIApp) ShowMainWindow() {
	if lapp.mainWindow == nil {
		return
	}
	lapp.mainWindow.Show()
}
