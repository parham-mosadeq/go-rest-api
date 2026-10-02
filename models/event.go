package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event

func (e Event) Save() {
	// * Implement the logic to save the event to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	// * Implement the logic to retrieve all events from the database
	return events
}
