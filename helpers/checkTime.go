package helpers

import (
	"fmt"
	"regexp"
)

func checkTime(time string) bool {
	pattern := `^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return reg.MatchString(time)
}
