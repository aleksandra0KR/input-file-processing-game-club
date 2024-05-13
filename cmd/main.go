package main

import (
	"bufio"
	"fmt"
	"inputfileprocess/helpers"
	"inputfileprocess/model"
	"inputfileprocess/parsers"
	"inputfileprocess/processors"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you should provide name of a input file")
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fileOutput, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		_ = fmt.Errorf("failed to create file %f: ", err)
		return
	}
	defer func(fileOutput *os.File) {
		err := fileOutput.Close()
		if err != nil {

		}
	}(fileOutput)

	// reading number of tables and check for valid input
	N := parsers.ParseNumber(scanner, "number of tables")

	// reading open and close time
	openTime, closeTime := parsers.ParseTime(scanner)
	line := fmt.Sprintf("%s\n", openTime.Format("15:04"))
	_, err = fileOutput.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	// reading cost of an hour
	cost := parsers.ParseNumber(scanner, "cost per hour")

	club := model.Club{
		Tables:         make(map[int]model.Table, N),
		Client:         make(map[string]model.Client),
		AmountOfTables: N,
		OpenTime:       *openTime,
		CloseTime:      *closeTime,
		PricePerHour:   cost,
		WaitingList:    make([]model.Client, 0)}

	// possessing events
	for scanner.Scan() {

		event := parsers.ParseEvent(scanner, club)
		switch event.EventID {
		case 1:
			processors.FirstEvent(event, &club, fileOutput)
		case 2:
			processors.SecondEvent(event, &club, fileOutput)
		case 3:
			processors.ThirdEvent(event, &club, fileOutput)
		case 4:
			processors.FourthEvent(event, &club, fileOutput)
		}

	}

	// removing all remaining clients
	for _, client := range club.Client {
		processors.EleventhEvent(&model.Event{ClientID: client.ClientID, TimeOfEvent: *closeTime}, &club, fileOutput)
	}

	line = fmt.Sprintf("%s\n", closeTime.Format("15:04"))
	_, err = fileOutput.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	// sort the slice by TableID
	tables := make([]model.Table, 0, len(club.Tables))
	for _, v := range club.Tables {
		tables = append(tables, v)
	}
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].TableID < tables[j].TableID
	})

	// print income
	for _, table := range tables {

		totalCoast := table.Payment * club.PricePerHour
		line = fmt.Sprintf("%d %d %s\n", table.TableID, totalCoast, helpers.DurationFormat(table.Exploitation))
		_, err = fileOutput.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
