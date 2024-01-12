package controllers

import (
	"github.com/alihaamedi/go-backend-events/models"
	"github.com/alihaamedi/go-backend-events/utility"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		failed400(ctx)
		return
	}

	err = user.Save()

	if err != nil {
		failed500(ctx)
		return
	}

	ok201(ctx, user)
}

func LogIn(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		failed400(ctx)
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		failed400(ctx)
		return
	}
	token, err := utility.GenerateToken(user.Email, user.ID)

	if err != nil {
		failed400(ctx)
		return
	}
	ok200(ctx, token)
}
