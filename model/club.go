package model

import "time"

type Club struct {
	TablesID       map[int]bool
	UsersID        map[string]int
	AmountOfTables int
	OpenTime       time.Time
	CloseTime      time.Time
	PricePerHour   int
}
