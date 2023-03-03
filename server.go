package main

import (
	"fmt"
	"golinebot/api"

	docs "golinebot/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	lineCtrl *api.LineController
	router   *gin.Engine
}

func NewServer(lineCtrl *api.LineController) *Server {
	return &Server{
		lineCtrl: lineCtrl,
		router:   gin.Default(),
	}
}

func (server *Server) SetupRouter() {
	server.router.POST("/api/v1/linebot/callback", server.lineCtrl.Callback)
	server.router.POST("/api/v1/linebot/message", server.lineCtrl.PushMsg)
	server.router.GET("/api/v1/linebot/:lineuserId/message", server.lineCtrl.Messages)
	docs.SwaggerInfo.BasePath = "/api/v1"
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (server *Server) Run(port int) {
	server.router.Run(fmt.Sprintf(":%d", port))
}
