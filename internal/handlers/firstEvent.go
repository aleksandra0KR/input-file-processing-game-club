package handlers

import (
	"fmt"
	model2 "inputfileprocess/internal/models"
	"os"
)

func FirstEvent(event *model2.Event, club *model2.Club, file *os.File) {

	line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}
	club.HistoryList = append(club.HistoryList, event.TimeOfEvent)

	if event.TimeOfEvent.Before(club.OpenTime) || event.TimeOfEvent.After(club.CloseTime) {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "NotOpenYet")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	} else {
		if _, ok := club.Client[event.ClientID]; ok {
			line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "YouShallNotPass")
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}
		}

		club.Client[event.ClientID] = model2.Client{
			ClientID:    event.ClientID,
			ArrivalTime: event.TimeOfEvent}
	}

}
