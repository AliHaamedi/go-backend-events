package controllers

import (
	"net/http"
	"strconv"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/alihaamedi/go-backend-events/res"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		res.Failed500(ctx)
		return
	}
	res.Ok200(ctx, events)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		res.Failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		res.Failed500(ctx)
		return
	}
	res.Ok201(ctx, event)
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		res.Failed400(ctx)
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		res.Failed404(ctx)
		return
	}
	res.Ok200(ctx, event)
}

func UpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		res.Failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		res.Failed404(ctx)
		return
	}

	if event.UserID != userId {
		res.Failed403(ctx)
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		res.Failed500(ctx)
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		res.Failed500(ctx)
		return
	}
	res.Ok200(ctx, updatedEvent)
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		res.Failed400(ctx)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		res.Failed404(ctx)
		return
	}

	if event.UserID != userId {
		res.Failed403(ctx)
		return
	}

	err = event.Delete()
	if err != nil {
		res.Failed500(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully!"})
}
