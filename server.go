package main

import (
	"fmt"
	"golinebot/api"

	"github.com/gin-gonic/gin"
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
}

func (server *Server) Run(port int) {
	server.router.Run(fmt.Sprintf(":%d", port))
}
