//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"golinebot/api"
	"golinebot/config"
	"golinebot/service"
)

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
	ConfigProvider,
	service.NewLineBot,
	api.NewLineController,
	NewServer,
)

func initapp() (*Server, func(), error) {
	panic(wire.Build(providerSet))
}
