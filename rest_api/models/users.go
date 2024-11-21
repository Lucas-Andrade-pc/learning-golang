package models

import "restapi/db"

type Users struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (users Users) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.PointerSqlBd.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(users.Email, users.Password) // insert the values at query
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	users.Id = id
	defer stmt.Close()
	return err
}
