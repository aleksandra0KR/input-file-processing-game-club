package model

import "time"

type Event struct {
	TimeOfEvent time.Time
	EventID     int
	ClientName  string
	TableID     int
}
