package ui

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type SystemTray struct {
	lapp			*LaunchdUIApp
	mSystemTrayIcon	*widgets.QSystemTrayIcon
	menu 			*widgets.QMenu
}

func NewSystemTray(app * LaunchdUIApp) *SystemTray {
	st := SystemTray{
		lapp:app,
		mSystemTrayIcon: widgets.NewQSystemTrayIcon(nil),
		menu:widgets.NewQMenu(nil),
	}
	st.init()
	return &st
}

func (st *SystemTray) init() {
	icon := gui.NewQIcon5(":/qml/tray.png")
	st.mSystemTrayIcon.SetIcon(icon)
	st.menu.AddAction("Exit").ConnectTriggered(func(bool) {
		st.lapp.Exit()
	})
	st.menu.AddAction("Show MainWindow").ConnectTriggered(func(bool) {
		st.lapp.ShowMainWindow()
	})
	st.mSystemTrayIcon.SetContextMenu(st.menu)
}

func (st *SystemTray) Launch() {
	st.mSystemTrayIcon.Show()
}

func (st *SystemTray) Close() {
	if st.mSystemTrayIcon == nil {
		return
	}
	st.mSystemTrayIcon.SetVisible(false)
}