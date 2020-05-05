package controller

import (
	"github.com/yiwenlong/launchduidemo/helper"
	"path/filepath"
)

type MacOSServerController struct {
	appRootDirPath	string
}

func (mctl *MacOSServerController) IsStart() bool {
	return false
}

func (mctl *MacOSServerController) Start(handler helper.ProcessHandler) {
	executable := filepath.Join(mctl.appRootDirPath, "server")
	err := helper.ConfigServerOnMacOS(mctl.appRootDirPath, executable)
	if err != nil {
		handler.HandleEcho(SessionStart, "ERROR: "+err.Error())
		return
	}
	startSh := filepath.Join(mctl.appRootDirPath, "boot")
	helper.ExecShellAsync(startSh, handler, SessionStart)
}

func (mctl *MacOSServerController) Stop(handler helper.ProcessHandler) {

}

func (mctl *MacOSServerController) LogFilePath() string {
	return ""
}