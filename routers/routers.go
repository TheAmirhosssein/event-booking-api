package routers

import "github.com/gin-gonic/gin"

func RegisterRouters(server *gin.Engine) {
	server.GET("/events", eventsHandler)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
}
