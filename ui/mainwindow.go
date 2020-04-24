package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/yiwenlong/launchduidemo/controller"
	"github.com/yiwenlong/launchduidemo/shell"
)

type MainWindow struct {
	app           *LaunchdUIApp
	window        *widgets.QMainWindow
	centralWidget *widgets.QWidget
	loggerWidget  *widgets.QTextEdit
	btnStart      *widgets.QPushButton
	btnClose      *widgets.QPushButton
	btnShowLog    *widgets.QPushButton
	servCtl       *controller.ServerController
	dialog        *Dialog
}

type Dialog struct {
	widgets.QMessageBox

	_ func(message string) `slot:"info"`
}

func NewMainWindow(app *LaunchdUIApp) *MainWindow {
	window := MainWindow{
		app:           app,
		window:        widgets.NewQMainWindow(nil, 0),
		centralWidget: widgets.NewQWidget(nil, 0),
		loggerWidget:  widgets.NewQTextEdit(nil),
		btnStart:      widgets.NewQPushButton2("start launchd test server", nil),
		btnClose:      widgets.NewQPushButton2("close launchd test server", nil),
		btnShowLog:    widgets.NewQPushButton2("show log", nil),
		servCtl:       controller.NewServerController(app.AppRootDir()),
		dialog:        NewDialog(nil),
	}
	window.init()
	return &window
}

func (mw *MainWindow) init() {
	mw.window.SetMinimumSize2(800, 400)
	mw.window.SetWindowTitle("launchd demo")

	mw.centralWidget.SetLayout(widgets.NewQVBoxLayout())
	mw.centralWidget.Layout().QLayoutItem.SetAlignment(core.Qt__AlignLeft)
	mw.window.SetCentralWidget(mw.centralWidget)

	mw.btnStart.ConnectClicked(func(bool) {
		mw.servCtl.Start(mw)
	})
	mw.centralWidget.Layout().AddWidget(mw.btnStart)

	mw.btnClose.ConnectClicked(func(bool) {
		mw.servCtl.Stop(mw)
	})
	mw.centralWidget.Layout().AddWidget(mw.btnClose)

	mw.loggerWidget.SetReadOnly(true)
	mw.centralWidget.Layout().AddWidget(mw.loggerWidget)

	mw.dialog.ConnectInfo(func(message string) {
		mw.dialog.Information(nil, "Infomation", message, widgets.QMessageBox__Ok, widgets.QMessageBox__Default)
	})
}

func (mw *MainWindow) HandleEcho(_ shell.SessionToken, echo string) {
	mw.loggerWidget.Append(echo)
}

func (mw *MainWindow) HandleSuccess(token shell.SessionToken) {
	switch token {
	case controller.Start:
		mw.app.app.SetQuitOnLastWindowClosed(false)
		mw.app.LaunchSystemTray()
		mw.dialog.Info("Service Started!")
	case controller.Stop:
		mw.app.app.SetQuitOnLastWindowClosed(true)
		mw.app.CloseSystemTray()
		mw.dialog.Info("Service Stop!")
	}
}

func (mw *MainWindow) HandleError(_ shell.SessionToken, _ int, state string) {
	mw.loggerWidget.Append("[ Shell exec error ]" + state)
}

func (mw *MainWindow) Launch() {
	mw.Show()
	if mw.servCtl.IsStarted() {
		mw.app.LaunchSystemTray()
	}
}

func (mw *MainWindow) Show() {
	mw.window.Show()
}

func (mw *MainWindow) Close() {
	mw.window.SetVisible(false)
}
