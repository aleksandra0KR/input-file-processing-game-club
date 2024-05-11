package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func EleventhEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 11, event.ClientID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	client, _ := club.Client[event.ClientID]
	client.DepartureTime = event.TimeOfEvent
	club.HistoryList = append(club.HistoryList, client)
	delete(club.Client, client.ClientID)
}
