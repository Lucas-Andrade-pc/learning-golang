package routes

import (
	"fmt"
	"net/http"
	events "restapi/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse param to int."})
		return
	}
	_, err = events.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cloud not fetch event"})
		return
	}

	var upEvent events.Event
	err = context.ShouldBindJSON(&upEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse request data."})
		return
	}

	fmt.Println(eventId)
	upEvent.Id = eventId
	e, err := upEvent.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not fetch update event"})
		return
	}
	context.JSON(http.StatusOK, e)
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse param to int."})
		return
	}
	event, err := events.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cloud not fetch event"})
		return
	}

	eventDelete, err := event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cloud not delete event"})
		return
	}
	context.JSON(http.StatusOK, eventDelete)

}

func getContainer(context *gin.Context) {
	addrs := events.Container()
	context.JSON(http.StatusOK, addrs[2])
}
