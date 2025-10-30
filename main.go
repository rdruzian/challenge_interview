package main

import (
	"fmt"

	"github.com/rdruzian/challenge_interview/docs"
	"github.com/rdruzian/challenge_interview/server"
)

// @title           Challenge code
// @version         1.0
// @description     API for challenge code interview, that I need to develop a CRUD system
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// Configura metadados do Swagger em runtime (garante consistência)
	docs.SwaggerInfo.Title = "challenge_interview"
	docs.SwaggerInfo.Description = "API for challenge code interview, that I need to develop a CRUD syste"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Create an instance of server
	server := server.NewServer()
	// Initiate a server
	server.Run()

	fmt.Println("Server running on 8080")
}
