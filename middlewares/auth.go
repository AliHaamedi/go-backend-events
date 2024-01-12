package middlewares

import (
	"net/http"

	"github.com/alihaamedi/go-backend-events/utility"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		notAuthorized(ctx)
		return
	}

	userId, err := utility.VerifyToken(token)
	if err != nil {
		notAuthorized(ctx)
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}

func notAuthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
}
