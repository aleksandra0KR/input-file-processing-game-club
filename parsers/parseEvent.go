package parsers

import (
	"bufio"
	"fmt"
	"inputfileprocess/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseEvent(scanner *bufio.Scanner) *model.Event {

	var event model.Event

	parts := strings.Split(scanner.Text(), " ")

	layout := "15:04"
	timeOfEvent, err := time.Parse(layout, parts[0])
	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	event.TimeOfEvent = timeOfEvent

	eventID, err := strconv.Atoi(parts[1])
	if err != nil || eventID <= 0 {
		fmt.Println("invalid ID of event:", parts[1])
		os.Exit(1)
	}

	event.EventID = eventID

	event.ClientID = parts[2]
	fmt.Println(parts[2])

	if eventID == 2 {
		tableID, err := strconv.Atoi(parts[3])
		if err != nil || tableID <= 0 {
			fmt.Println("invalid ID of table:", parts[3])
			os.Exit(1)
		}
		event.TableID = tableID
	}

	scanner.Scan()

	return &event
}
