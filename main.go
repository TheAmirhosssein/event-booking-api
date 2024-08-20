package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/events", eventsHandler)
	server.Run(":8080")
}

func eventsHandler(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Hello There!"})
}
