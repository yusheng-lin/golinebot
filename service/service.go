package service

import (
	"golinebot/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewLineBot(config *config.Config) (*linebot.Client, error) {
	return linebot.New(config.LineChannelSecret, config.LineChannelToken)
}
