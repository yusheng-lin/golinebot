package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	server, cleanup, err := initapp()
	defer cleanup()
	if err != nil {
		log.Fatal().Err(err)
		return
	}
	server.SetupRouter()
	server.Run(8080)
}
