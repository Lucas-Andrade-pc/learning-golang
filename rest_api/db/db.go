package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	connect, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Cloud not connect to database")
	}

	connect.SetMaxOpenConns(10)
	connect.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventTables := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTERGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL
	  user_id INTERGER
	)
	`
	_, err := DB.Exec(createEventTables)
	if err != nil {
		panic("cloud not create table ")
	}
}
