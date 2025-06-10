package configs

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App *App `yaml:"app"`
	DB  *DB  `yaml:"db"`
}

func (c *Config) Validate() error {
	if err := c.App.Validate(); err != nil {
		return fmt.Errorf(`failed to validate app config: %s`, err.Error())
	}
	if err := c.DB.Validate(); err != nil {
		return fmt.Errorf(`failed to validate db config: %s`, err.Error())
	}

	return nil
}

func LoadConfig(fileName string, config *Config) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf(`cant read file '%s'`, fileName)
	}

	if err = yaml.Unmarshal(file, config); err != nil {
		return fmt.Errorf(`file %s yaml unmarshal error: %s`, fileName, err.Error())
	}
	if config.App.Env == "dev" {
		fmt.Printf("current config: \n\n%s\n", pp.Sprint(config))
	}
	return config.Validate()
}
