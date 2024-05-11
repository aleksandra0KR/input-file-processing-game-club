package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParseNumber(scanner *bufio.Scanner) int {
	number, err := strconv.Atoi(scanner.Text())
	if err != nil || number < 0 {
		fmt.Println("invalid number of tables:", scanner.Text())
		os.Exit(1)
	}
	scanner.Scan()
	return number
}
