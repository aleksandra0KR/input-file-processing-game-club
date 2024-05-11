package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ParseTime(scanner *bufio.Scanner) (*time.Time, *time.Time) {
	parts := strings.Split(scanner.Text(), " ")

	layout := "15:04"
	openTime, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	closeTime, err := time.Parse(layout, parts[1])
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	scanner.Scan()

	return &openTime, &closeTime
}
