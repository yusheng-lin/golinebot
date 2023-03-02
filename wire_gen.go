// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"golinebot/api"
	"golinebot/config"
	"golinebot/service"
)

// Injectors from wire.go:

func initapp() (*Server, func(), error) {
	config, err := ConfigProvider()
	if err != nil {
		return nil, nil, err
	}
	client, err := service.NewLineBot(config)
	if err != nil {
		return nil, nil, err
	}
	lineController := api.NewLineController(client)
	server := NewServer(lineController)
	return server, func() {
	}, nil
}

// wire.go:

var cf *config.Config

func ConfigProvider() (*config.Config, error) {
	if cf == nil {
		c, err := config.NewConfig("./env")

		if err != nil {
			return c, err
		}
		cf = c
	}
	return cf, nil
}

var providerSet = wire.NewSet(
	ConfigProvider, service.NewLineBot, api.NewLineController, NewServer,
)