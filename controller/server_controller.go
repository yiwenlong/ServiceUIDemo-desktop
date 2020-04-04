package controller

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/yiwenlong/launchduidemo/shell"
)

type ServerController struct {
	homeDir 	*core.QDir
	startScript	string
	stopScript	string
}

const (
	Start shell.ShellToken = iota
	Stop
)

func NewServerController(serverHomeDir *core.QDir) *ServerController {
	sc := ServerController{
		homeDir:     serverHomeDir,
	}
	startSh := serverHomeDir.AbsoluteFilePath("start.sh")
	sc.startScript = fmt.Sprintf("%s %s", startSh, serverHomeDir.AbsolutePath())
	stopSh := serverHomeDir.AbsoluteFilePath("stop.sh")
	sc.stopScript = fmt.Sprintf("%s %s", stopSh, serverHomeDir.AbsolutePath())
	return  &sc
}

func (sc *ServerController) Start(handler shell.ShellHandler) {
	shell.ExecShellAsync(sc.startScript, handler, Start)
}

func (sc *ServerController) Stop(handler shell.ShellHandler) {
	shell.ExecShellAsync(sc.stopScript, handler, Stop)
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