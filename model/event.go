package model

import "time"

type event struct {
	timeOfEvent time.Time
	eventID     int8
	clientName  string
	tableID     int64
}
