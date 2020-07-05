package config

import "errors"

type DarwinConfig struct {
}

func (config *DarwinConfig) Config(homeDir, executable string, args ...string) error {
	return errors.New("Drawin config not implement current.")
}
