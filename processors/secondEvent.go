package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func SecondEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent, event.EventID, event.ClientName, event.TableID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}
	_, ok := club.UsersID[event.ClientName]
	if !ok {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 13, "ClientUnknown")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

	if club.TablesID[event.TableID] == false {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent, 13, "PlaceIsBusy")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

}
