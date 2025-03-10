package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var PointerSqlBd *sql.DB

func InitDb() {

	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		fmt.Println("Error Connect to database")
		panic("Cloud not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables(DB)

	PointerSqlBd = DB

}

// Function init tables databases
func createTables(db *sql.DB) {

	createUserTables := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := db.Exec(createUserTables)
	if err != nil {
		panic(fmt.Sprintf("Erro create users tables - %v", err))
	}
	createEventTables := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTERGER,
	  FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = db.Exec(createEventTables)
	if err != nil {
		panic(fmt.Sprintf("Erro create event tables - %v", err))
	}
	fmt.Println("Create tables sucesso!")
}
