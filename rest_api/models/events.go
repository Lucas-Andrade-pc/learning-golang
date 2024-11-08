package events

import (
	"database/sql"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DataTime    time.Time `binding:"required"`
	UserID      int
}

var event = []Event{}

func (e Event) Save() error {
	var db *sql.DB
	query := `INSERT INTO events(name, description, location, dataTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DataTime, e.UserID) // insert the values at query
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	defer stmt.Close()
	return err
}

func GetAllEvent() []Event {
	return event
}
