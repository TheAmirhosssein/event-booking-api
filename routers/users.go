package routers

import (
	"net/http"

	"github.com/TheAmirhosssein/event-booking-api/models"
	"github.com/TheAmirhosssein/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json"})
		return
	}
	err = user.Save()
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			context.JSON(http.StatusNotFound, gin.H{"message": "this username have taken"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json"})
		return
	}
	err = user.ValidateCredential()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	accessKey, err := utils.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"access_key": accessKey})
}
