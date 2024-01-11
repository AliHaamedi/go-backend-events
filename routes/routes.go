package routes

import (
	"github.com/alihaamedi/go-backend-events/controllers"
	"github.com/alihaamedi/go-backend-events/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	//events
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	//need authentication group
	authServer := server.Group("/")
	authServer.Use(middlewares.Authenticate)
	authServer.POST("/events", controllers.CreateEvent)
	authServer.PUT("/events/:id", controllers.UpdateEvent)
	authServer.DELETE("/events/:id", controllers.DeleteEvent)
	authServer.POST("/events/:id/register", controllers.RegisterForEvent)
	authServer.DELETE("/events/:id/register")

	//users
	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.LogIn)
}
