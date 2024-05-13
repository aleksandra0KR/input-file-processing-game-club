package helpers

func Contains(element int) bool {
	eventsId := []int{1, 2, 3, 4, 11, 12, 13}
	for _, e := range eventsId {
		if e == element {
			return true
		}
	}
	return false
}
