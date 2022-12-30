package config

import (
	"github.com/dmahmalat/black-friday-board-game/pkg/errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	App struct {
		Name    string `env-required:"true"`
		Version string `env-required:"true"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}
)

const (
	APPNAME       = "bfbg"
	CONFIGFILE    = "config.yaml"
	VERSION       = "0.0"
	VERSIONSTRING = "custom"
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.App.Name = APPNAME
	cfg.App.Version = VERSION + "." + VERSIONSTRING

	err := cleanenv.ReadConfig("./"+CONFIGFILE, cfg)
	errorCheck(err)

	err = cleanenv.ReadEnv(cfg)
	errorCheck(err)

	return cfg, nil
}

func errorCheck(e error) {
	errors.ErrorCheck(e, "config error:")
}
