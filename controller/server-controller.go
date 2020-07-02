package controller

import (
	"github.com/yiwenlong/launchduidemo/helper"
	"runtime"
)

type IServerController interface {
	IsStart() bool
	Start(helper.ProcessHandler)
	Stop(helper.ProcessHandler)
	LogFilePath() string
}

func NewServerController(appRootDirPath string) IServerController {
	if runtime.GOARCH == "drawin" {
		return &MacOSServerController{
			appRootDirPath: appRootDirPath,
		}
	} else {
		return &WindowsServerController{
			appRootDirPath: appRootDirPath,
		}
	}
}

const (
	SessionStart helper.SessionToken = iota
	SessionStop
)
