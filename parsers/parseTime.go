package parsers

import (
	"time"
)

func ParseTime(t string) (*time.Time, error) {
	layout := "15:04"
	openTime, err := time.Parse(layout, t)
	if err != nil {
		return nil, err
	}
	return &openTime, nil
}
