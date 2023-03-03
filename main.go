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
		log.Error().Err(err)
		return
	}
	log.Print("app start")
	server.SetupRouter()
	server.Run(8080)
}
