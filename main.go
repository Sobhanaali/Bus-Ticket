package main

import (
	"os"
	"github.com/Sobhanaali/Bus-Ticket/internal/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// load config

	config , err := util.LoadConfig(".")

	if(err != nil) {
		log.Fatal().Err(err).Msg("Cannot load config !")
	}

	if(config.APPDEBUG == "true") {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}