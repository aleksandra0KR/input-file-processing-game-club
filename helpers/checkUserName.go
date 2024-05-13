package helpers

import (
	"fmt"
	"regexp"
)

func CheckUserName(username string) bool {
	pattern := `^[a-zA-Z0-9_-]*$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return reg.MatchString(username)
}
