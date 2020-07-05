package config

type Helper interface {
	Config(home string, cmd string, args ...string) error
}
