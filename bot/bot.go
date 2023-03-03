package bot

import (
	"golinebot/config"
	"golinebot/model"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewLineBot(config *config.Config) (*linebot.Client, error) {
	return linebot.New(config.LineChannelSecret, config.LineChannelToken)
}

func ParseMessage(event *linebot.Event) *model.Message {
	return &model.Message{
		LineUserId: event.Source.UserID,
		Text:       event.Message.(*linebot.TextMessage).Text,
		Time:       event.Timestamp,
	}
}
