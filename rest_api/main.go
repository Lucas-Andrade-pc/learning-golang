package main

import (
	"fmt"
	"net/http"
	"restapi/db"
	events "restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)
	server.POST("/events", createEvent)
	server.GET("/nginx", getNginx)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	event, err := events.GetAllEvent()
	if err != nil {
		fmt.Println("entrei no if getEvents")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro return all events"})
	}
	context.JSON(http.StatusOK, event)
}
func getEventsById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse param to int."})
		return
	}

	res, err := events.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cloud not fetch event id"})
		return
	}
	context.JSON(http.StatusOK, res)

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

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not create event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event create!", "event": event})
}

func getNginx(context *gin.Context) {
	addrs := events.Nginx()
	context.JSON(http.StatusOK, addrs[2])
}
