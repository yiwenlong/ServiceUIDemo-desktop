package config

import "runtime"

type Helper interface {
	Config(home string, cmd string, args ...string) error
}

var ConfigHelper Helper

func init() {
	switch runtime.GOOS {
	case "darwin":
		ConfigHelper = &DarwinConfig{}
	case "windows":
		ConfigHelper = &WindowsConfig{}
	default:
		panic("Operation system not support: " + runtime.GOOS)
	}
}
