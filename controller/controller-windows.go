package controller

import (
	"github.com/yiwenlong/launchduidemo/helper"
	"path/filepath"
)

type WindowsServerController struct {
	appRootDirPath	string
}

func (wctl *WindowsServerController) IsStart() bool {
	return false
}

func (wctl *WindowsServerController) Start(handler helper.ProcessHandler) {
	executable := filepath.Join(wctl.appRootDirPath, "server")
	err := helper.ConfigServerOnWindows(wctl.appRootDirPath, executable)
	if err != nil {
		handler.HandleEcho(SessionStart, "ERROR: "+err.Error())
		return
	}
	startSh := filepath.Join(wctl.appRootDirPath, "boot")
	helper.ExecShellAsync(startSh, handler, SessionStart)
}

func (wctl *WindowsServerController) Stop(handler helper.ProcessHandler) {
	stopSh := filepath.Join(wctl.appRootDirPath, "stop")
	helper.ExecShellAsync(stopSh, handler, SessionStop)
}

func (wctl *WindowsServerController) LogFilePath() string {
	return ""
}