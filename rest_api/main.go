package main

import (
	"net/http"
	"restapi/db"
	events "restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	event := events.GetAllEvent()
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event events.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse request data."})
		return
	}
	event.Id = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event create!", "event": event})
}
