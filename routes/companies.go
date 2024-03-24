package routes

import (
	"encoding/json"

	"example.com/digital-passport/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var requestData map[string]any
	err := json.NewDecoder(context.Request.Body).Decode(&requestData)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse company request data"})
		return
	}

	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not convert company request data to JSON"})
		return
	}

	var company models.Company
	err = json.Unmarshal(requestJSON, &company)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse company request data"})
		return
	}

	var user models.User
	err = json.Unmarshal(requestJSON, &user)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse user request data"})
		return
	}

	if company.CompanyName == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		context.JSON(500, gin.H{"message": "fields missing"})
		return
	}

	companyId, err := company.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save company"})
		return
	}

	user.Company = companyId
	userId, err := user.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(201, gin.H{"message": "user created successfully", "companyId": companyId, "userId": userId})
}
