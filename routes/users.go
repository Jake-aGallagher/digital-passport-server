package routes

import (
	"fmt"

	"example.com/digital-passport/models"
	"github.com/gin-gonic/gin"
)

func signupUser(context *gin.Context) {
	var user models.User
	fmt.Println(context.Request.Body)
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}
	fmt.Println("company: ", user)

	insertId, err := user.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(201, gin.H{"message": "user created successfully", "id": insertId})
}
