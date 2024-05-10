package helpers

import "time"

func checkSubsequenceOfTimeStamps(firstEvent, secondEvent time.Time) bool {
	if firstEvent.Before(secondEvent) {
		return false
	}
	return true
}
