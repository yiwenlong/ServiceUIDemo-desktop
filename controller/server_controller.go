package controller

import (
	"github.com/therecipe/qt/core"
	"github.com/yiwenlong/launchduidemo/helper"
	"github.com/yiwenlong/launchduidemo/shell"
)

type IServerController interface {
	IsStart() bool
	Start()
	Stop()
	LogFilePath() string
}

func NewIServerController() *IServerController{
	return nil
}

type ServerController struct {
	homeDir *core.QDir
	mShell  shell.Shell
}

const (
	Start shell.SessionToken = iota
	Stop
)

func (sc *ServerController) Start(handler shell.CommandHandler) {
	executable := sc.homeDir.AbsoluteFilePath("server")
	err := helper.ConfigServer(sc.homeDir.AbsolutePath(), executable)
	if err != nil {
		handler.HandleEcho(Start, "ERROR: "+err.Error())
		return
	}
	startSh := sc.homeDir.AbsoluteFilePath("boot")
	sc.mShell.ExecAsync(startSh, []string{sc.homeDir.AbsolutePath()}, handler, Start)
}

func (sc *ServerController) Stop(handler shell.CommandHandler) {
	stopSh := sc.homeDir.AbsoluteFilePath("stop")
	sc.mShell.ExecAsync(stopSh, []string{sc.homeDir.AbsolutePath()}, handler, Stop)
}

func (sc *ServerController) LogFile() *core.QFile {
	return core.NewQFile2(sc.homeDir.AbsoluteFilePath("error.log"))
}

func (sc *ServerController) Log() string {
	logFile := sc.LogFile()
	logFile.Open(core.QIODevice__ReadOnly)
	log := logFile.ReadAll()
	logFile.Close()
	return log.Data()
}

func (sc *ServerController) IsStarted() bool {
	// state, _ := shell.ExecShell("launchctl list | grep \"com.1wenlong.server\"")
	// return state.Success()
	return false
}
