package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	OrderTopic string   `yaml:"order_topic"`
	Brokers    []string `yaml:"brokers"`
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
