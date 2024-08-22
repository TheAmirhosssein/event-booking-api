package models

import (
	"github.com/TheAmirhosssein/event-booking-api/db"
)

func RegisterEvent(userId, eventId int64) error {
	query := `
	INSERT INTO registration (user_id, event_id)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, eventId)
	return err
}

func DeleteRegistration(userId, eventId int64) error {
	query := `DELETE From registration where user_id = ? AND event_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, eventId)
	return err
}
