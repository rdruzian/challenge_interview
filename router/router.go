package router

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/rdruzian/challenge_interview/inbound"
	"github.com/rdruzian/challenge_interview/middleware"
	"github.com/rdruzian/challenge_interview/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var deviceService inbound.DeviceInterface

func ConfigRoutes(router *gin.Engine, service inbound.DeviceInterface) *gin.Engine {
	deviceService = service

	router.Use(middleware.ErrorHandler())
	router.NoRoute(middleware.NotFoundHandler)

	v1 := router.Group("/api/v1")
	{
		// Rotas para produtos
		device := v1.Group("/device")
		{
			device.POST("/create", createDevice)
			device.GET("", getAllDevices)
			device.GET("/:id", getDeviceById)
			device.GET("/brand/:brand", getDevicesByBrand)
			device.GET("/state/:state", getDevicesByState)
			device.DELETE("/:id", deleteDevice)
		}
	}

    // Rota para documentação Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// @Summary      Create a new device
// @Description  Create a new device on database
// @Tags         device
// @Accept       json
// @Produce      json
// @Param        device body model.Device true "Device"
// @Success      201  {object}  model.Device
// @Failure      400  {object}  map[string]string
// @Router       /device/create [post]
func createDevice(c *gin.Context) {
	var device model.Device

	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Debug("request_device", "device", device)

	if err := deviceService.CreateDevice(device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": device})
}

// @Summary      Get all devices
// @Tags         device
// @Produce      json
// @Success      200  {array}  model.Device
// @Failure      500  {object}  map[string]string
// @Router       /device [get]
func getAllDevices(c *gin.Context) {
	devices, err := deviceService.GetAllDevice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// @Summary      Get device by ID
// @Tags         device
// @Produce      json
// @Param        id path int true "Device ID"
// @Success      200  {object}  model.Device
// @Failure      400  {object}  map[string]string
// @Router       /device/{id} [get]
func getDeviceById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	device, err := deviceService.GetDevice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// @Summary      Get device by Brand
// Description   Get all devices by brand
// @Tags         device
// @Produce      json
// @Param        brand path string true "Brand"
// @Success      200  {array}  model.Device
// @Failure      500  {object}  map[string]string
// @Router       /device/brand/{brand} [get]
func getDevicesByBrand(c *gin.Context) {
	brand := c.Param("brand")
	devices, err := deviceService.GetDeviceByBrand(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// @Summary      Get device by State
// Description   Get all devices by state
// @Tags         device
// @Produce      json
// @Param        state path string true "State"
// @Success      200  {array}  model.Device
// @Failure      500  {object}  map[string]string
// @Router       /device/state/{state} [get]
func getDevicesByState(c *gin.Context) {
	state := c.Param("state")
	devices, err := deviceService.GetDeviceByState(state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// @Summary      Delete a single device
// @Tags         device
// @Produce      json
// @Param        id path int true "Device ID"
// @Success      204
// @Failure      400  {object}  map[string]string
// @Router       /device/{id} [delete]
func deleteDevice(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := deviceService.DeleteDevice(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
