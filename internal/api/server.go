package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func InitServer() *Server {
	e := gin.New()

	return &Server{
		Engine: e,
	}
}

func (e *Server) StartServer(port string) {
	if err := e.Engine.Run(port); err != nil {
		log.Fatalf("error on run server: %s", err)
	}
}
