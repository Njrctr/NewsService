package configs

import "errors"

type App struct {
	HttpPort       int    `yaml:"httpPort"`
	LogQueries     bool   `yaml:"logQueries"`
	GinReleaseMode bool   `yaml:"ginReleaseMode"`
	Env            string `yaml:"env"`
}

func (c *App) Validate() error {
	if c.HttpPort == 0 {
		return errors.New(`httpPort is zero`)
	}

	return nil
}
