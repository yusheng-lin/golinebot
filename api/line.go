package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/rs/zerolog/log"
)

type LineController struct {
	linebot *linebot.Client
}

func NewLineController(linebot *linebot.Client) *LineController {
	return &LineController{
		linebot: linebot,
	}
}

func (ctrl *LineController) Callback(ctx *gin.Context) {
	events, err := ctrl.linebot.ParseRequest(ctx.Request)

	if err != nil {
		log.Error().Err(err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			fmt.Println(event.Message)
		}
	}
}
