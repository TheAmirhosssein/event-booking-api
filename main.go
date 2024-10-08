package main

import (
	"github.com/TheAmirhosssein/event-booking-api/db"
	"github.com/TheAmirhosssein/event-booking-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()
	routers.RegisterRouters(server)
	server.Run("0.0.0.0:8080")
}
