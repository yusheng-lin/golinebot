package api

import (
	"golinebot/bot"
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
			msg := bot.ParseMessage(event)
			if err := ctrl.svc.AddMsg(msg); err != nil {
				ctx.JSON(http.StatusInternalServerError, struct{}{})
			}
		}
	}
}

// @Summary push message
// @Tags linebot
// @version 1.0
// @produce application/json
// @param lineuserId path string true "lineuserId"
// @param text body model.Text true "msg"
// @Success 200 string string 成功後返回的值
// @Router /linebot/{lineuserId}/message [post]
func (ctrl *LineController) PushMsg(ctx *gin.Context) {
	lineuserId := ctx.Param("lineuserId")
	msg := &model.Text{}
	err := ctx.BindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusOK, "Invalid params")
		return
	}
	_, err = ctrl.linebot.PushMessage(lineuserId, linebot.NewTextMessage(msg.Text)).Do()

	if err != nil {
		log.Error().Err(err).Msg("")
		ctx.JSON(http.StatusInternalServerError, "push message fail please try again later")
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

// @Summary fetch messages
// @Tags linebot
// @version 1.0
// @produce application/json
// @param lineuserId path string true "lineuserId"
// @Success 200 string string 成功後返回的值
// @Router /linebot/{lineuserId}/message [get]
func (ctrl *LineController) Messages(ctx *gin.Context) {
	lineuserId := ctx.Param("lineuserId")
	msgs, err := ctrl.svc.GetMessages(lineuserId)

	if err != nil {
		log.Error().Err(err).Msg("")
		ctx.JSON(http.StatusInternalServerError, "fetch message fail please try again later")
		return
	}
	ctx.JSON(http.StatusOK, msgs)
}
