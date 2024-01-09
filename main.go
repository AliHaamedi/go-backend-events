package main

import (
	"github.com/alihaamedi/go-backend-events/db"
	"github.com/alihaamedi/go-backend-events/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoute(server)
	server.Run(":8080") // localhost:8080
}
