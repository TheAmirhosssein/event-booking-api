package main

import (
	"net/http"

	"github.com/TheAmirhosssein/event-booking-api/models"
	"github.com/gin-gonic/gin"
	"github.com/TheAmirhosssein/event-booking-api/db"
	
)

func main() {
	server := gin.Default()
	db.InitDB()
	server.GET("/events", eventsHandler)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func eventsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	var incomingData models.Event
	err := context.BindJSON(&incomingData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not parse json"})
		return
	}
	incomingData.ID = 1
	incomingData.Save()
	context.JSON(200, incomingData)
}
