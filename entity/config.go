package entity

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// CONFIG_LOCATION is the default localtion
// of configuration file. This file is located at
// $HOME/.gostrap.yaml
const CONFIG_LOCATION = "%s/.gostrap.yaml"

// Configuration struct is the type
// that contains GoStrap configuration variable.
type Configuration struct {
	Template            string `mapstructure:"template"`
	GoModulePlaceholder string `mapstructure:"gomoduleplaceholder"`
}

// Update function is use to update $HOME/.gostrap.yaml
// configuration value.
func (c *Configuration) Update() error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	osUserDir, _ := os.UserHomeDir()
	return ioutil.WriteFile(fmt.Sprintf(CONFIG_LOCATION, osUserDir), out, 0755)
}
