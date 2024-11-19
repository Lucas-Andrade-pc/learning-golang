package events

import (
	"fmt"
	"net"
	"restapi/db"
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
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.PointerSqlBd.Prepare(query)
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
func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.PointerSqlBd.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &event, nil
}
func GetAllEvent() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.PointerSqlBd.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var events []Event

	for rows.Next() { //rows.Next se ainda houver mais linhas ele retorna true
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func Nginx() []net.Addr {
	ifaces, _ := net.Interfaces()

	var address []net.Addr
	for _, v := range ifaces {
		addrs, _ := v.Addrs()
		for n, _ := range addrs {
			address = append(address, addrs[n])
		}

	}
	fmt.Println(address)
	return address
}
