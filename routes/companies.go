package routes

import (
	"fmt"

	"example.com/digital-passport/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var company models.Company
	fmt.Println(context.Request.Body)
	err := context.ShouldBindJSON(&company)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}
	fmt.Println("company: ", company)

	insertId, err := company.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(201, gin.H{"message": "user created successfully", "id": insertId})
}
