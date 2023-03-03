package service

import (
	"golinebot/model"

	"github.com/rs/zerolog/log"
)

type IRepository interface {
	AddMessage(msg *model.Message) error
	GetMessages(lineuserId string) (*[]model.Message, error)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) AddMsg(msg *model.Message) error {
	err := svc.repo.AddMessage(msg)

	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	return nil
}

func (svc *Service) GetMessages(lineuserId string) (*[]model.Message, error) {
	msgs, err := svc.repo.GetMessages(lineuserId)

	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	return msgs, nil
}
