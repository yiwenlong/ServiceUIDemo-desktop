package config

type DarwinConfig struct {
}

func (config *DarwinConfig) Config(homeDir, executable string, args ...string) error {
	return nil
}
