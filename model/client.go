package model

import "time"

type Client struct {
	ClientID      int
	Table         *Table
	ArrivalTime   time.Time
	DepartureTime time.Time
}
