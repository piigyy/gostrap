package entity

import (
	"os"

	"gopkg.in/yaml.v2"
)

const CONFIG_LOCATION = "$HOME/.gostrap.yaml"

type Configuration struct {
	Template            string `mapstructure:"template"`
	GoModulePlaceholder string `mapstructure:"goModulePlaceholder"`
}

func (c *Configuration) Update() error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, _ := os.Create(CONFIG_LOCATION)
	f.WriteString(string(out))
	if err != nil {
		return err
	}

	defer f.Close()
	return nil
}
