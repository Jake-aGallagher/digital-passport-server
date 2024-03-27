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

	server.GET("/passports/:companyid", middlewares.Authenticate, getPassports) // view all passports relevant to a company
	server.GET("/passports/:companyid/:id", middlewares.Authenticate)           // view a single passport
	server.POST("/passports", middlewares.Authenticate, createPassport)         // creating a brand new passport
	server.POST("/passport-stages", middlewares.Authenticate)                   // adding a new stage to the passport
	server.PUT("/passport-stages/:id", middlewares.Authenticate)                // updating a stage to the passport
	server.DELETE("passport-stages/:id", middlewares.Authenticate)              // deleting a stage, only if no stages after this one

	server.POST("/files")
	server.GET("/files/passport/:id")
	server.GET("/files/file/:id")
	server.DELETE("/files/:id")

	server.POST("/signup-company", signup)
	server.POST("/signup-user", signupUser)
	server.POST("/login", Login)
	server.POST("/logout")
}
