package routes

import (
	"fmt"

	"example.com/digital-passport/models"
	"github.com/gin-gonic/gin"
)

func getPassports(context *gin.Context) {
	companyId := context.Param("companyid")
	if companyId == "" {
		context.JSON(400, gin.H{"message": "Could not parse company id"})
		return
	}
	fmt.Println("company id: ", companyId)
	passports, err := models.GetPassportsForCompany(companyId)
	if err != nil {
		context.JSON(500, gin.H{"message": "error retrieving passports"})
		return
	}

	context.JSON(200, gin.H{"message": "passports retrieved successfully", "passports": passports})
}

func createPassport(context *gin.Context) {
	companyId, exists := context.Get("companyId")
	if !exists {
		context.JSON(400, gin.H{"message": "error retrieving company id"})
		return
	}

	var passport models.Passport
	err := context.ShouldBindJSON(&passport)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}
	passport.CompanyId = companyId.(string)

	passport.Save()
}
