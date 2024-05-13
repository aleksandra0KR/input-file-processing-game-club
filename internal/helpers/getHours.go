package helpers

import (
	"time"
)

func GetHours(d time.Duration) int {
	mints := -int(d.Minutes())
	hrs := mints / 60
	if mints%60 > 0 {
		hrs++
	}
	return hrs
}
