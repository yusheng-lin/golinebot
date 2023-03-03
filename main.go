package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
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
