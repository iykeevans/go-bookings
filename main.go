package main

import (
	"github.com/gin-gonic/gin"
	"go-bookings.com/db"
	"go-bookings.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
