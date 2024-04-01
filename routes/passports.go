package routes

import (
	"fmt"
	"strings"
	"time"

	"example.com/digital-passport/models"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-diceware/diceware"
)

func getPassports(context *gin.Context) {
	companyId := context.Param("companyid")
	if companyId == "" {
		context.JSON(400, gin.H{"message": "Could not parse company id"})
		return
	}

	passports, err := models.GetPassportsForCompany(companyId)
	if err != nil {
		context.JSON(500, gin.H{"message": "error retrieving passports"})
		return
	}

	context.JSON(200, gin.H{"message": "passports retrieved successfully", "passports": passports})
}

func getPassport(context *gin.Context) {
	companyId := context.Param("companyid")
	passportId := context.Param("id")
	if companyId == "" || passportId == "" {
		context.JSON(400, gin.H{"message": "Could not parse request"})
		return
	}
	passport, err := models.GetPassportById(companyId, passportId)
	if err != nil {
		context.JSON(500, gin.H{"message": "error retrieving passport"})
		return
	}
	context.JSON(200, gin.H{"message": "passport retrieved successfully", "passport": passport})
}

func addPassport(context *gin.Context) {
	companyId, exists := context.Get("companyId")
	if !exists {
		context.JSON(400, gin.H{"message": "error retrieving company id"})
		return
	}

	var passport models.Passport
	passport.CompanyId = companyId.(string)
	passport.Files = []string{"1234"}
	passport.Locked = false
	passport.Created = time.Now()
	err := context.ShouldBindJSON(&passport)
	if err != nil {
		fmt.Println("some err: ", err)
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}

	if passport.Locked {
		list, err := diceware.Generate(4)
		if err != nil {
			context.JSON(500, gin.H{"message": "unable to generate passcode"})
			return
		}
		passport.UseCode = strings.Join(list, "-")
	}

	_, err = passport.Save()
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not save passport"})
		return
	}
	context.JSON(200, gin.H{"message": "passport saved successfully"})
}

func editPassport(context *gin.Context) {
	companyId, exists := context.Get("companyId")
	if !exists {
		context.JSON(400, gin.H{"message": "error retrieving company id"})
		return
	}

	passportId := context.Param("id")
	if passportId == "" {
		context.JSON(400, gin.H{"message": "Could not parse passport id"})
		return
	}

	foundPassport, err := models.GetPassportById(companyId.(string), passportId)
	if err != nil || foundPassport.Locked {
		context.JSON(500, gin.H{"message": "passport already locked"})
		return
	}
	var passport models.Passport = foundPassport
	err = context.ShouldBindJSON(&passport)
	if err != nil {
		fmt.Println("some err: ", err)
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}

	if passport.Locked {
		list, err := diceware.Generate(4)
		if err != nil {
			context.JSON(500, gin.H{"message": "unable to generate passcode"})
			return
		}
		passport.UseCode = strings.Join(list, "-")
	}

	_, err = passport.Save()
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not save passport"})
		return
	}
	context.JSON(200, gin.H{"message": "passport saved successfully"})
}
