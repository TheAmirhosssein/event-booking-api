package routers

import (
	"net/http"
	"strconv"

	"github.com/TheAmirhosssein/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func eventsHandler(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var incomingData models.Event
	userId := context.GetInt64("userId")
	incomingData.UserID = &userId
	err := context.BindJSON(&incomingData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not parse json"})
		return
	}
	incomingData.ID = 1
	err = incomingData.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(201, incomingData)
}

func getEvent(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
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
	userId := context.GetInt64("userId")
	if *event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you can not update this event"})
		return
	}
	var updatedEvent models.Event
	err = context.BindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not parse json"})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context) {
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
	userId := context.GetInt64("userId")
	if *event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you can not delete this event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
