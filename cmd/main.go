package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You should provide name of a input file")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		_ = fmt.Errorf("error occured while opening the file %f", err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("error occured while closing the file %f", err)
			os.Exit(1)
		}
	}(file)

	// reading number of tables and check for valid input
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
	N, err := strconv.Atoi(scanner.Text())
	if err != nil || N <= 0 {
		fmt.Println("invalid number of tables:", scanner.Text())
		return
	}
	scanner.Scan()
	fmt.Println(N)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
