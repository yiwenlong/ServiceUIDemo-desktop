package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/yiwenlong/ServiceUIDemo-desktop/controller"
	"github.com/yiwenlong/ServiceUIDemo-desktop/helper"
)

type MainWindow struct {
	app           *LaunchdUIApp
	window        *widgets.QMainWindow
	centralWidget *widgets.QWidget
	loggerWidget  *widgets.QTextEdit
	btnStart      *widgets.QPushButton
	btnClose      *widgets.QPushButton
	btnShowLog    *widgets.QPushButton
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
		btnStart:      widgets.NewQPushButton2("start server", nil),
		btnClose:      widgets.NewQPushButton2("stop server", nil),
		btnShowLog:    widgets.NewQPushButton2("show log", nil),
		dialog:        NewDialog(nil),
	}
	window.init()
	return &window
}

func (mw *MainWindow) init() {
	mw.window.SetMinimumSize2(800, 400)
	mw.window.SetWindowTitle("Service UI Demo")

	mw.centralWidget.SetLayout(widgets.NewQVBoxLayout())
	mw.centralWidget.Layout().QLayoutItem.SetAlignment(core.Qt__AlignLeft)
	mw.window.SetCentralWidget(mw.centralWidget)

	mw.btnStart.ConnectClicked(func(bool) {
		mw.app.serverCtl.Start(mw)
	})
	mw.centralWidget.Layout().AddWidget(mw.btnStart)

	mw.btnClose.ConnectClicked(func(bool) {
		mw.app.serverCtl.Stop(mw)
	})
	mw.centralWidget.Layout().AddWidget(mw.btnClose)

	mw.loggerWidget.SetReadOnly(true)
	mw.centralWidget.Layout().AddWidget(mw.loggerWidget)

	mw.dialog.ConnectInfo(func(message string) {
		mw.dialog.Information(nil, "Information", message, widgets.QMessageBox__Ok, widgets.QMessageBox__Default)
	})
}

func (mw *MainWindow) Launch() {
	mw.Show()
	if mw.app.serverCtl.IsStart() {
		mw.app.LaunchSystemTray()
	}
}

func (mw *MainWindow) Show() {
	mw.window.Show()
}

func (mw *MainWindow) Close() {
	mw.window.SetVisible(false)
}

//// MARK: Implement ProcessCallback Interface methods.
func (mw *MainWindow) Echo(_ helper.SessionToken, echo string) {
	mw.loggerWidget.Append(echo)
}

func (mw *MainWindow) OnSuccess(token helper.SessionToken) {
	switch token {
	case controller.SessionStart:
		mw.app.app.SetQuitOnLastWindowClosed(false)
		mw.app.LaunchSystemTray()
		mw.dialog.Info("Service Started!")
	case controller.SessionStop:
		mw.app.app.SetQuitOnLastWindowClosed(true)
		mw.app.CloseSystemTray()
		mw.dialog.Info("Service SessionStop!")
	}
}

func (mw *MainWindow) OnError(_ helper.SessionToken, _ int, state string) {
	mw.loggerWidget.Append("[ Shell exec error ]" + state)
}
