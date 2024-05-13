package handlers

import (
	"fmt"
	model2 "inputfileprocess/internal/models"
	"os"
)

func ThirdEvent(event *model2.Event, club *model2.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}
	club.HistoryList = append(club.HistoryList, event.TimeOfEvent)

	if club.AmountOfTables-len(club.Tables) > 0 {
		line := fmt.Sprintf("%s %d %s %s\n", event.TimeOfEvent.Format("15:04"), 13, event.ClientID, "ICanWaitNoLonger")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	} else if len(club.WaitingList) > club.AmountOfTables {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 11, event.ClientID)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	} else {
		club.WaitingList = append(club.WaitingList, model2.Client{ClientID: event.ClientID, ArrivalTime: event.TimeOfEvent})
	}

}
