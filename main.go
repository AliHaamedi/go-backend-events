package main

import (
	"net/http"

	"github.com/alihaamedi/go-backend-events/db"
	"github.com/alihaamedi/go-backend-events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)    // localhost:8080/events
	server.POST("/events", createEvent) // localhost:8080/events

	server.Run(":8080") // localhost:8080
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world!", "data": events})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "all fields are required"})
		return
	}
	event.UserId = 1
	eventId, err := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not crate event"})
		return
	}
	event.ID = eventId

	ctx.JSON(http.StatusCreated, gin.H{"message": "event was created successfully", "data": event})
}
