package main

import (
	"fmt"

	"github.com/rdruzian/challenge_interview/server"
)

// @title           Challenge code
// @version         1.0
// @description     API for challenge code interview, that I need to develop a CRUD system
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// Create an instance of server
	server := server.NewServer()
	// Initiate a server
	server.Run()

	fmt.Println("Server running on 8080")
}
