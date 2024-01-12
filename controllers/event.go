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
		failed500(ctx)
		return
	}
	ok200(ctx, events)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		failed500(ctx)
		return
	}
	ok201(ctx, event)
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		failed400(ctx)
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		failed404(ctx)
		return
	}
	ok200(ctx, event)
}

func UpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		failed404(ctx)
		return
	}

	if event.UserID != userId {
		failed403(ctx)
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		failed500(ctx)
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		failed500(ctx)
		return
	}
	ok200(ctx, updatedEvent)
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		failed404(ctx)
		return
	}

	if event.UserID != userId {
		failed403(ctx)
		return
	}

	err = event.Delete()
	if err != nil {
		failed500(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully!"})
}
