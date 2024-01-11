package controllers

import (
	"net/http"
	"strconv"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func UnRegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.UnRegister(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not un register user for event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "unRegistered successfully"})
}
