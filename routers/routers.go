package routers

import (
	"github.com/TheAmirhosssein/event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {
	server.GET("/events", middlewares.AuthenticateMiddleware, eventsHandler)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/sign-up", signUp)
	server.POST("/login", login)
}
