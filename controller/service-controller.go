package controller

import (
	"fmt"
	"github.com/yiwenlong/ServiceUIDemo-desktop/controller/config"
	"github.com/yiwenlong/ServiceUIDemo-desktop/helper"
	"path/filepath"
)

type ServiceController interface {
	Start() (chan string, error)
	Stop() (chan string, error)
	IsStart() bool
	ShowLog()
}

type ServiceControllerImpl struct {
	ServiceHome string
	ProcessName string
}

func (sc *ServiceControllerImpl) Start() (chan string, error) {

}

func (sc *ServiceControllerImpl) Stop() (chan string, error) {

}

func (sc *ServiceControllerImpl) IsStart() bool {

}

func (sc *ServiceControllerImpl) ShowLog() bool {

}


type ServerController struct {
	config     config.Helper
	serverHome string
}

func (servCtl *ServerController) Start(callback helper.ProcessCallback) {
	executable := filepath.Join(servCtl.serverHome, "server")
	if err := servCtl.config.Config(servCtl.serverHome, executable); err != nil {
		callback.Echo(SessionStart, "ERROR: "+err.Error())
		return
	}
	startSh := filepath.Join(servCtl.serverHome, "boot")
	helper.ExecShellAsync(fmt.Sprintf("%s %s", startSh, servCtl.serverHome), callback, SessionStart)
}

func (servCtl *ServerController) Stop(callback helper.ProcessCallback) {
	stopSh := filepath.Join(servCtl.serverHome, "stop")
	helper.ExecShellAsync(stopSh, callback, SessionStop)
}

func (servCtl *ServerController) LogFilePath() string {
	return filepath.Join(servCtl.serverHome, "server.log")
}

func (servCtl *ServerController) IsStart() bool {
	return false
}

func NewServerController(home string) *ServerController {
	return &ServerController{
		config:     config.ConfigHelper,
		serverHome: home,
	}
}

const (
	SessionStart helper.SessionToken = iota
	SessionStop
)
