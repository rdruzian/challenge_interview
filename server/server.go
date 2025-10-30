package server

import (
	"log"
	"os"

	"github.com/rdruzian/challenge_interview/database"
	"github.com/rdruzian/challenge_interview/repository"
	"github.com/rdruzian/challenge_interview/router"
	"github.com/rdruzian/challenge_interview/service"

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
	// Modo somente Swagger: pula DB e serviços para facilitar preview/local
	if os.Getenv("SWAG_ONLY") == "true" {
		r := router.ConfigRoutes(s.server, nil)
		log.Fatal(r.Run(":" + s.port))
		return
	}

	// Initialize database
	database.StartDB()
	db := database.GetDatabase()

	// Wire repository and service
	repo := repository.NewDeviceRepository(db)
	deviceService := service.NewDeviceService(repo)

	// Configure routes with service and start server
	r := router.ConfigRoutes(s.server, deviceService)
	log.Fatal(r.Run(":" + s.port))
}
