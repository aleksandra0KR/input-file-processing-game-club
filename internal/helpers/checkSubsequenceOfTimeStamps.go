package helpers

import "time"

func CheckSubsequenceOfTimeStamps(firstEvent, secondEvent time.Time) bool {
	if firstEvent.After(secondEvent) {
		return false
	}
	return true
}
