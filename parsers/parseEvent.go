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

	if len(parts) < 3 {
		fmt.Println("error input: ", scanner.Text())
		os.Exit(1)
	}

	layout := "15:04"
	timeOfEvent, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("error parsing time: ", err, scanner.Text())
		os.Exit(1)
	}
	if len(club.HistoryList) > 0 && !helpers.CheckSubsequenceOfTimeStamps(club.HistoryList[len(club.HistoryList)-1], timeOfEvent) {
		fmt.Println("Error subsequence of time stamps")
		os.Exit(1)
	}
	club.HistoryList = append(club.HistoryList, timeOfEvent)

	event.TimeOfEvent = timeOfEvent

	eventID, err := strconv.Atoi(parts[1])
	if err != nil || eventID <= 0 {
		fmt.Println("invalid ID of event:", scanner.Text())
		os.Exit(1)
	}
	if !helpers.Contains(eventID) {
		fmt.Println("invalid ID of event:", scanner.Text())
		os.Exit(1)
	}

	event.EventID = eventID

	if !helpers.CheckUserName(parts[2]) {
		fmt.Println("invalid user name", scanner.Text())
		os.Exit(1)
	}
	event.ClientID = parts[2]

	if eventID == 2 {
		if len(parts) != 4 {
			fmt.Println("error input: ", scanner.Text())
			os.Exit(1)
		}
		tableID, err := strconv.Atoi(parts[3])
		if err != nil || tableID <= 0 {
			fmt.Println("invalid ID of table:", scanner.Text())
			os.Exit(1)
		}
		if !helpers.CheckTableNumber(tableID, club.AmountOfTables) {
			fmt.Println("invalid table number", scanner.Text())
			os.Exit(1)
		}
		event.TableID = tableID
	}

	if len(parts) > 3 && eventID != 2 {
		fmt.Println("error input: ", scanner.Text())
		os.Exit(1)
	}
	return &event
}
