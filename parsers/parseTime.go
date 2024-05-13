package parsers

import (
	"bufio"
	"fmt"
	"inputfileprocess/helpers"
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
	if helpers.CheckTime(parts[0]) {
		fmt.Println("Invalid format of time")
		os.Exit(1)
	}
	openTime, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	closeTime, err := time.Parse(layout, parts[1])
	if helpers.CheckTime(parts[1]) {
		fmt.Println("Invalid format of time")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	return &openTime, &closeTime
}
