package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func ThirdEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent, event.EventID, event.ClientName, event.TableID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	var availableTables []int
	for i, table := range club.TablesID {
		if table == true {
			availableTables = append(availableTables, i)
		}
	}

	if len(availableTables) > 0 {
		line := fmt.Sprintf("%s %d %s %s\n", event.TimeOfEvent, 13, event.ClientName, "ICanWaitNoLonger")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

	if len(club.WaitingList) >= club.AmountOfTables {
		line := fmt.Sprintf("%s %d %s %s\n", event.TimeOfEvent, 11, event.ClientName)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

}
