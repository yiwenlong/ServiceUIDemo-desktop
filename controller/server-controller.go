package controller

import (
	"github.com/yiwenlong/launchduidemo/helper"
)

type IServerController interface {
	IsStart() bool
	Start(helper.ProcessHandler)
	Stop(helper.ProcessHandler)
	LogFilePath() string
}

func NewIServerController(appRootDirPath string) IServerController{
	ctl := MacOSServerController{
		appRootDirPath: appRootDirPath,
	}
	return &ctl
}

const (
	SessionStart helper.SessionToken = iota
	SessionStop
)

