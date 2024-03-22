package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/passports/:companyid")     // view all passports relevant to a company
	server.GET("/passports/:companyid/:id") // view a single passport
	server.POST("/passports")               // creating a brand new passport
	server.POST("/passport-stages")         // adding a new stage to the passport
	server.PUT("/passport-stages/:id")      // updating a stage to the passport
	server.DELETE("passport-stages/:id")    // deleting a stage, only if no stages after this one

	server.POST("/files")
	server.GET("/files/passport/:id")
	server.GET("/files/file/:id")
	server.DELETE("/files/:id")

	server.POST("/signup-company")
	server.POST("/signup-user")
	server.POST("/login")
	server.POST("/logout")
}
