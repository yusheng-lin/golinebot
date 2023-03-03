//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golinebot/api"
	"golinebot/bot"
	"golinebot/config"
	"golinebot/db"
	"golinebot/service"
	"time"
)

var cf *config.Config
var mgdb *mongo.Client

func configProvider() (*config.Config, error) {
	if cf == nil {
		c, err := config.NewConfig("./env")

		if err != nil {
			return c, err
		}
		cf = c
	}
	return cf, nil
}

func mongodbProvider(config *config.Config) (*mongo.Client, error) {
	if mgdb == nil {
		conn := fmt.Sprintf("mongodb://%s:%s@%s", config.MongoUser, config.MongoPwd, config.MongoURL)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))

		if err != nil {
			return nil, err
		}

		mgdb = client
	}
	return mgdb, nil
}

var providerSet = wire.NewSet(
	configProvider,
	mongodbProvider,
	db.NewRepository,
	bot.NewLineBot,
	service.NewService,
	api.NewLineController,
	NewServer,
)

func initapp() (*Server, func(), error) {
	panic(wire.Build(providerSet))
}
