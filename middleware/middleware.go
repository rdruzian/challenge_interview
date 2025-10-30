package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for error treatment
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Loga o erro
			for _, e := range c.Errors {
				log.Printf("Erro: %v", e.Err)
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error to process your request",
			})
		}
	}
}

// NotFoundHandler is responsible for not found routes
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Route not found",
	})
}
