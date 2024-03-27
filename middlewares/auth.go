package middlewares

import (
	"example.com/digital-passport/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorised"})
		return
	}

	userId, companyId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorised"})
		return
	}

	context.Set("userId", userId)
	context.Set("companyId", companyId)

	context.Next()
}
