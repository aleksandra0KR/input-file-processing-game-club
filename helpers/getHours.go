package helpers

import (
	"time"
)

func GetHours(d time.Duration) int {
	mins := -int(d.Minutes())
	hrs := mins / 60
	if mins%60 > 0 {
		hrs++
	}
	return hrs
}
