package parsers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ParseNumber(scanner *bufio.Scanner, value string) int {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}

	number, err := strconv.Atoi(scanner.Text())
	if err != nil || number < 0 {
		fmt.Println("invalid ", value, " ", scanner.Text())
		os.Exit(1)
	}
	return number
}
