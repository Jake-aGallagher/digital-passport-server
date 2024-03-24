package routes

import (
	"fmt"

	"example.com/digital-passport/models"
	"example.com/digital-passport/utils"
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

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}

	companyId, userId, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(401, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(companyId, userId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(200, gin.H{"message": "Login successfull", "token": token, "companyId": companyId, "userId": userId})
}
