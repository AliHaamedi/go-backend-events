package routes

import (
	"github.com/alihaamedi/go-backend-events/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events/:id", controllers.GetEvent)
}
