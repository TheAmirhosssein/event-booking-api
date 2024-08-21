package main

import (
	"net/http"
	"strconv"

	"github.com/TheAmirhosssein/event-booking-api/db"
	"github.com/TheAmirhosssein/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()
	server.GET("/events", eventsHandler)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)
	server.Run("127.0.0.1:8080")
}

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
	err := context.BindJSON(&incomingData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not parse json"})
		return
	}
	incomingData.ID = 1
	err = incomingData.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(200, incomingData)
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
