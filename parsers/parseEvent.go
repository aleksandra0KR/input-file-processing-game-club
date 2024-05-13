package parsers

import (
	"bufio"
	"fmt"
	"inputfileprocess/helpers"
	"inputfileprocess/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseEvent(scanner *bufio.Scanner, club model.Club) *model.Event {

	var event model.Event

	parts := strings.Split(scanner.Text(), " ")

	layout := "15:04"
	if !helpers.CheckTime(parts[0]) {
		fmt.Println("Error not required time format")
		os.Exit(1)
	}
	timeOfEvent, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}
	if !helpers.CheckSubsequenceOfTimeStamps(club.HistoryList[len(club.HistoryList)-1], timeOfEvent) {
		fmt.Println("Error subsequence of time stamps")
		os.Exit(1)
	}

	event.TimeOfEvent = timeOfEvent

	eventID, err := strconv.Atoi(parts[1])
	if err != nil || eventID <= 0 {
		fmt.Println("invalid ID of event:", parts[1])
		os.Exit(1)
	}
	if !helpers.Contains(eventID) {
		fmt.Println("Invalid event ID")
		os.Exit(1)
	}

	event.EventID = eventID

	if !helpers.CheckUserName(parts[2]) {
		fmt.Println("Invalid user name")
		os.Exit(1)
	}
	event.ClientID = parts[2]
	fmt.Println(parts[2])

	if eventID == 2 {
		tableID, err := strconv.Atoi(parts[3])
		if err != nil || tableID <= 0 {
			fmt.Println("invalid ID of table:", parts[3])
			os.Exit(1)
		}
		if !helpers.CheckTableNumber(tableID, club.AmountOfTables) {
			fmt.Println("Invalid table number")
		}
		event.TableID = tableID
	}
	return &event
}
