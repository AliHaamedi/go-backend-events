package controllers

import (
	"net/http"

	"github.com/alihaamedi/go-backend-events/models"
	"github.com/alihaamedi/go-backend-events/utility"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "all fields are required"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not crate user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user was created successfully", "data": user})
}

func LogIn(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "all fields are required"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utility.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login is successful", "token": token})
}
