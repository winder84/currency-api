package config

import (
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DB  DB  `yaml:"db"`
	Api Api `yaml:"api"`
	App App `yaml:"app"`
}
type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Source string `yaml:"source"`
}

type Process struct {
	Ticker int `yaml:"ticker"`
}

type Api struct {
	AppId   string  `yaml:"app_id"`
	Url     string  `yaml:"url"`
	Process Process `yaml:"process"`
}

func New(filePath string) (config Config, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return config, errors.Wrapf(err, "open config file: %s", filePath)
	}
	defer func(err error) {
		multierr.AppendInto(&err, file.Close())
	}(err)

	yamlDecoder := yaml.NewDecoder(file)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return config, errors.Wrap(err, "decode config failed")
	}
	return config, nil
}
