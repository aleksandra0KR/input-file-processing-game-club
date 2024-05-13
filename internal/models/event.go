package models

import "time"

type Event struct {
	TimeOfEvent time.Time
	EventID     int
	ClientID    string
	TableID     int
}
