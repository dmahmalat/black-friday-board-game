package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/dmahmalat/black-friday-board-game/pkg/errors"
	"github.com/dmahmalat/black-friday-board-game/pkg/util"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		HTTP `yaml:"http"`
	}

	App struct {
		Name    string
		Version string
	}

	HTTP struct {
		Port int `yaml:"port" env:"HTTP_PORT"`
	}
)

const (
	APPNAME       string = "bfbg"
	CONFIGFILE    string = "config.yaml"
	VERSION       string = "0.0"
	VERSIONSTRING string = "custom"

	// flag - default values
	DEFAULT_HTTP_PORT int = 80
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{}

	// Set App name and version
	cfg.App.Name = APPNAME
	cfg.App.Version = VERSION + "." + VERSIONSTRING

	// Set defaults
	cfg.HTTP.Port = DEFAULT_HTTP_PORT

	// Read config file
	if util.FileExists("./" + CONFIGFILE) {
		err = cleanenv.ReadConfig("./"+CONFIGFILE, cfg)
		errorCheck(err)
	}

	// Read env vars
	err = cleanenv.ReadEnv(cfg)
	errorCheck(err)

	// Read CLI arguments
	err = processArgs(cfg)
	errorCheck(err)

	return cfg, err
}

// handle CLI arguments
func processArgs(cfg *Config) error {
	var err error
	f := flag.NewFlagSet(cfg.App.Name, flag.ExitOnError)

	// Port
	var port util.FlagVar
	f.Var(&port, "p", "HTTP `port` - default "+strconv.Itoa(DEFAULT_HTTP_PORT))

	// -help message
	fUsage_original := f.Usage
	f.Usage = func() {
		fUsage_original()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output(), "\n"+envHelp)
		fmt.Fprintln(f.Output(), "\nOrder of precedence:\ndefault values < config.yaml file < environment vars < command-line arguments")
	}

	// Read values from CLI
	f.Parse(os.Args[1:])

	// Set values after parsing
	if port.IsSet() {
		cfg.HTTP.Port, err = strconv.Atoi(port.String())
	}

	return err
}

func errorCheck(e error) {
	errors.ErrorCheck(e, "config error:")
}
