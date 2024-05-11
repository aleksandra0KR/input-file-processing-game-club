package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func FirstEvent(event *model.Event, club *model.Club, file *os.File) {

	line := fmt.Sprintf("%s %d %d\n", event.TimeOfEvent, event.EventID, event.ClientID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}

	if event.TimeOfEvent.Before(club.OpenTime) || event.TimeOfEvent.After(club.CloseTime) {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 13, "NotOpenYet")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	} else {
		if _, ok := club.Client[event.ClientID]; ok {
			line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 13, "YouShallNotPass")
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}
		}
	}

}
