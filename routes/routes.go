package routes

import (
	"github.com/alihaamedi/go-backend-events/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	//events
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events/:id", controllers.GetEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)

	//users
	server.POST("/signup", controllers.SignUp)
}
