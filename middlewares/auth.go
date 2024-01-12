package middlewares

import (
	"github.com/alihaamedi/go-backend-events/controllers"
	"github.com/alihaamedi/go-backend-events/utility"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		controllers.NotAuthorized(ctx)
		return
	}

	userId, err := utility.VerifyToken(token)
	if err != nil {
		controllers.NotAuthorized(ctx)
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
