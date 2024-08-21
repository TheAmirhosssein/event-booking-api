package models

import "github.com/TheAmirhosssein/event-booking-api/db"

type User struct {
	ID       int64
	Username string
	Password string
}

func (user *User) Save() error {
	query := `
	INSERT INTO users (username, password) 
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}
