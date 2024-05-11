package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func FourthEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	client, ok := club.Client[event.ClientID]
	if !ok {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "ClientUnknown")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	} else {
		if len(club.WaitingList) > 0 {
			nextClient := club.WaitingList[0]
			club.WaitingList = club.WaitingList[1:]

			nextClient.Table = client.Table
			nextClient.Table.Client = &nextClient
			delete(club.Client, client.ClientID)
			line = fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent.Format("15:04"), 12, nextClient.ClientID, nextClient.Table.TableID)
			_, err = file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}

		} else {
			delete(club.Client, client.ClientID)
		}

	}
}
