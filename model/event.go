package model

import "time"

type Event struct {
	TimeOfEvent time.Time
	EventID     int
	ClientID    int
	TableID     int
}
