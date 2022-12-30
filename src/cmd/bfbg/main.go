package main

import (
	"log"
	"os"

	"github.com/dmahmalat/black-friday-board-game/config"
	"github.com/dmahmalat/black-friday-board-game/pkg/errors"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	errorCheck(err)

	// Run
	log.Println(cfg.App.Name)
	log.Println(cfg.App.Version)
	log.Println(cfg.HTTP.Port)
	os.Exit(0)
}

func errorCheck(e error) {
	errors.ErrorCheck(e, "error at startup:")
}
