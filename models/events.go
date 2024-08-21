package models

import (
	"time"

	"github.com/TheAmirhosssein/event-booking-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      *int64
}

func (event *Event) Save() error {
	sqlCommand := `
	INSERT INTO events (name, description, location, dateTime)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime)
	if err != nil {
		return err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = insertedId
	return nil
}

func (event Event) Update() error {
	query := `
	UPDATE EVENTS
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	stmt.Exec(event.ID)
	defer stmt.Close()
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
