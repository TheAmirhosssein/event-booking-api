package models

import (
	"errors"

	"github.com/TheAmirhosssein/event-booking-api/db"
	"github.com/TheAmirhosssein/event-booking-api/utils"
)

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
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Username, password)
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

func (user *User) ValidateCredential() error {
	query := "SELECT password FROM users where username=?"
	row := db.DB.QueryRow(query, user.Username)
	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}
	passwordValidation := utils.CheckPasswordHash(user.Password, hashedPassword)
	if !passwordValidation {
		return errors.New("credentials invalid")
	}
	return nil
}
