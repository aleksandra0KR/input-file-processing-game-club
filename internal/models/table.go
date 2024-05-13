package models

import "time"

type Table struct {
	TableID             int
	Client              *Client
	StartOfExploitation time.Time
	EndOfExploitation   time.Time
	Exploitation        time.Duration
	Payment             int
}
