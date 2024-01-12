package controllers

import (
	"strconv"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.Register(userId)
	if err != nil {
		failed500(ctx)
		return
	}
	ok201(ctx, nil)
}

func UnRegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.UnRegister(userId)
	if err != nil {
		failed500(ctx)
		return
	}
	ok200(ctx, nil)
}
