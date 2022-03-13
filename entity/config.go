package entity

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const CONFIG_LOCATION = "%s/.gostrap.yaml"

type Configuration struct {
	Template            string `mapstructure:"template"`
	GoModulePlaceholder string `mapstructure:"gomoduleplaceholder"`
}

func (c *Configuration) Update() error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	osUserDir, _ := os.UserHomeDir()
	return ioutil.WriteFile(fmt.Sprintf(CONFIG_LOCATION, osUserDir), out, 0755)
}
