package parsers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ParseTime(scanner *bufio.Scanner) (*time.Time, *time.Time) {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	parts := strings.Split(scanner.Text(), " ")

	layout := "15:04"
	openTime, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("error parsing time:", err, scanner.Text())
		os.Exit(1)
	}

	closeTime, err := time.Parse(layout, parts[1])

	if err != nil {
		fmt.Println("error parsing time:", err, scanner.Text())
		os.Exit(1)
	}

	return &openTime, &closeTime
}
