package models

import "time"

type Client struct {
	ClientID      string
	Table         *Table
	ArrivalTime   time.Time
	DepartureTime time.Time
}
