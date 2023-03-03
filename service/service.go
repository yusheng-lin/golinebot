package service

import (
	"golinebot/config"
	"golinebot/db"
	"golinebot/model"

	"github.com/rs/zerolog/log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewLineBot(config *config.Config) (*linebot.Client, error) {
	return linebot.New(config.LineChannelSecret, config.LineChannelToken)
}

type Service struct {
	repo *db.Repository
}

func NewService(repo *db.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) AddMsg(msg *model.Receive) error {
	err := svc.repo.AddMessage(msg)

	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	return nil
}
