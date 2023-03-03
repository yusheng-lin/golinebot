package api

import (
	"golinebot/model"
	"golinebot/service"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineController struct {
	linebot *linebot.Client
	svc     *service.Service
}

func NewLineController(linebot *linebot.Client, svc *service.Service) *LineController {
	return &LineController{
		linebot: linebot,
		svc:     svc,
	}
}

func (ctrl *LineController) Callback(ctx *gin.Context) {
	events, err := ctrl.linebot.ParseRequest(ctx.Request)

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			msg := parseMessage(event)
			if err := ctrl.svc.AddMsg(msg); err != nil {
				ctx.JSON(http.StatusExpectationFailed, struct{}{})
			}
		}
	}
}

func parseMessage(event *linebot.Event) *model.Message {
	return &model.Message{
		LineUserId: event.Source.UserID,
		Text:       event.Message.(*linebot.TextMessage).Text,
		Time:       event.Timestamp,
	}
}
