package routers

import (
	"github.com/TheAmirhosssein/event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {
	server.GET("/events", eventsHandler)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.AuthenticateMiddleware)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)

	server.POST("/sign-up", signUp)
	server.POST("/login", login)
}
