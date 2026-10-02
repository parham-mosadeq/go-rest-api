package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event

func (e Event) Save() error {
	// * Implement the logic to save the event to the database
	query := `INSERT INTO event(name, description,location,dateTime,user_id) 
	VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() []Event {
	// * Implement the logic to retrieve all events from the database
	return events
}
