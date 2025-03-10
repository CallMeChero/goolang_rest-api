package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	// later: add to DB
	events = append(events, e)
}

func GetAllEvent() []Event {
	return events
}
