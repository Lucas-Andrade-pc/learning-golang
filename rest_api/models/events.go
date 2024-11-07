package events

import (
	"fmt"
	"time"
)

type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DataTime    time.Time `binding:"required"`
	UserID      int
}

var event = []Event{}

func (e Event) Save() {
	fmt.Println(e)
	event = append(event, e)
}

func GetAllEvent() []Event {
	return event
}
