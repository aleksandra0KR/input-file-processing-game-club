package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func FourthEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent, event.EventID, event.ClientID, event.TableID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	client, ok := club.Client[event.ClientID]
	if !ok {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 13, "ClientUnknown")
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

			line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 11, event.ClientID)
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}

			line = fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent, 12, nextClient.ClientID, nextClient.Table.TableID)
			_, err = file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}

		} else {
			line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 11, event.ClientID)
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}
		}

	}
}
