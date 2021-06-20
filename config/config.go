package config

import (
	"io/ioutil"
	"os"

	"github.com/galenguyer/retina/core"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Services []core.Service `yaml:"services"`
}

func Load(path string) (*Config, error) {
	return loadConfigFile(path)
}

func loadConfigFile(path string) (*Config, error) {
	if bytes, err := ioutil.ReadFile(path); err == nil {
		yamlBytes := []byte(os.ExpandEnv(string(bytes)))
		config := Config{}
		err = yaml.Unmarshal(yamlBytes, &config)
		if err != nil {
			return nil, err
		}
		return &config, nil
	}
	return nil, nil
}
