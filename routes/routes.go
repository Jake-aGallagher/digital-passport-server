package routes

import (
	"example.com/digital-passport/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	server.Use(cors.New(config))

	server.GET("/passports/:companyid", middlewares.Authenticate, getPassports)  // view all passports relevant to a company
	server.GET("/passports/passport/:id", middlewares.Authenticate, getPassport) // view a single passport
	server.POST("/passports", middlewares.Authenticate, addPassport)             // creating a brand new passport / stage
	server.PUT("/passports/:id", middlewares.Authenticate, editPassport)         // editing a passport stage
	server.DELETE("passport-stages/:id", middlewares.Authenticate)               // deleting an unlocked stage

	server.POST("/files", middlewares.Authenticate, addFile)    // adding a file to a passport stage
	server.GET("/files/:id", middlewares.Authenticate, getFile) // getting an individual file for downloading
	server.DELETE("/files/:id")                                 // deleting an individual file

	server.POST("/signup-company", signup)
	server.POST("/signup-user", signupUser)
	server.POST("/login", Login)
}
