package routers

import (
	"net/http"
	"strconv"

	"github.com/TheAmirhosssein/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid endpoint"})
		return
	}
	event, err := models.GetEvent(id)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "event does not exist"})
		return
	}
	err = models.RegisterEvent(userId, event.ID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "event does not exist"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func cancelEventRegistration(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid endpoint"})
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "event does not exist"})
		return
	}
	err = models.DeleteRegistration(context.GetInt64("userId"), event.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event does not exist"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successful"})

}
