package router

import (
	"supertaltest/server"

	"github.com/gin-gonic/gin"
)

func NewRoute(router *gin.Engine, handler *server.ApiHandler) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/managers/parking-lots", handler.ParkingManagerHandler.CreateParkingLot)
		v1.GET("/managers/parking-lots/:id", handler.ParkingManagerHandler.GetParkingSlotStatus)
		v1.GET("/managers/parking-summaries", handler.ParkingManagerHandler.ParkingSummary)
		v1.PUT("/managers/parking-slots/:id/maintenance", handler.ParkingManagerHandler.ToggleParkingSlotMaintenance)
		v1.POST("/parking-lots/:id/park", handler.ParkingHandler.ParkUserVehicle)
		v1.POST("/tickets/:code/exit", handler.ParkingHandler.ExitUserVehicle)
	}
}
