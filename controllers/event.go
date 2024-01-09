package controllers

import (
	"net/http"
	"strconv"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world!", "data": events})
}

func CreateEvent(ctx *gin.Context) {
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

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "event could not be found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": event})
}
