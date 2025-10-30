package router

import (
	"TestInterview/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	// Adiciona middleware de tratamento de erros
	router.Use(middleware.ErrorHandler())

	// Configura handler para rotas não encontradas
	router.NoRoute(middleware.NotFoundHandler)

	// Grupo de rotas para a API v1
	v1 := router.Group("/api/v1")
	{
		// Rotas para produtos
		device := v1.Group("/device")
		{
			device.POST("/create", createDevice)

		}
	}

	// Rota para documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor na porta 8080
	log.Println("Server running on 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start serve: %v", err)
	}

	return router
}

// @Summary      Create a new device
// @Description  Create a new device on database
// @Tags         device
// @Produce      json
// @Param        device
// @Success      200  {object}  models.Product
// @Failure      404  {object}  map[string]string
// @Router       /device/create [post]
func createDevice(c *gin.Context) {

}
