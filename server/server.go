package server

import (
	"github.com/rdruzian/challenge_interview/router"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	r := router.ConfigRoutes(s.server)

	log.Fatal(r.Run(":" + s.port))
}
