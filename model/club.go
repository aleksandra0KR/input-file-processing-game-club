package model

import "time"

type Club struct {
	Tables         map[int]Table
	Client         map[int]Client
	AmountOfTables int
	OpenTime       time.Time
	CloseTime      time.Time
	PricePerHour   int
	WaitingList    []Client
}
