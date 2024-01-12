package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func res(message string, data any) gin.H {
	return gin.H{"message": message, "data": data}

}

// StatusOK
func Ok200(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, res("done!", data))
}

// StatusCreated
func Ok201(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, res("created...", data))
}

// StatusBadRequest
func Failed400(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, res("bad request", nil))
}

// StatusForbidden
func Failed403(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, res("Forbidden", nil))
}

// StatusNotFound
func Failed404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, res("could not found", nil))
}

// StatusInternalServerError
func Failed500(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, res("something went wrong", nil))
}

// StatusUnauthorized
func NotAuthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res("not authorized", nil))
}
