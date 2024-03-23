package main

import (
	"example.com/digital-passport/db"
	"example.com/digital-passport/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
