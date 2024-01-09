package main

import (
	"net/http"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)    // localhost:8080/events
	server.POST("/events", createEvent) // localhost:8080/events

	server.Run(":8080") // localhost:8080
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world!", "data": events})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "all fields are required"})
		return
	}
	event.ID = 1
	event.UserId = 1
	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "event was created successfully", "data": event})
}
