package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ServicesConfig struct {
	Database string `yaml:"database"`
}

type ConfigStruct struct {
	Services ServicesConfig `yaml:"services"`
	Port     int32          `omitempty,yaml:"port"`
}

var ConfigData ConfigStruct

func Init() error {
	rawConfig, err := os.ReadFile("config.yml")

	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawConfig, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing config file")
	}

	return nil
}
