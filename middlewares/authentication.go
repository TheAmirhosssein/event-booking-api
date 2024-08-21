package middlewares

import (
	"net/http"
	"strings"

	"github.com/TheAmirhosssein/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		context.Abort()
		return
	}
	if len(strings.Split(token, " ")) != 2 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		context.Abort()
		return
	}
	token = strings.Split(token, " ")[1]
	claims, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	context.Set("user", claims)
	context.Next()
}
