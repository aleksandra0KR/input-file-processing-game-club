package helpers

import (
	"fmt"
	"time"
)

func DurationFormat(d time.Duration) string {
	d = -1 * d
	secs := int(d.Seconds())
	hours := secs / 3600
	secs %= 3600
	mints := secs / 60
	secs %= 60
	return fmt.Sprintf("%02d:%02d", hours, mints)
}
