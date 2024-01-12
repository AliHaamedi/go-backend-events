package controllers

import (
	"strconv"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/alihaamedi/go-backend-events/res"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.Register(userId)
	if err != nil {
		res.Failed500(ctx)
		return
	}
	res.Ok200(ctx, nil)
}

func UnRegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.UnRegister(userId)
	if err != nil {
		res.Failed500(ctx)
		return
	}
	res.Ok200(ctx, nil)
}
