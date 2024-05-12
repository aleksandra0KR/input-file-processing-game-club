package processors

import (
	"fmt"
	"inputfileprocess/model"
	"os"
)

func SecondEvent(event *model.Event, club *model.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID, event.TableID)
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
	}

	if club.Tables[event.TableID].Client != nil {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "PlaceIsBusy")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
	}

	table, ok := club.Tables[event.TableID]
	if !ok {
		club.Tables[event.TableID] = model.Table{
			TableID:             event.TableID,
			Client:              &client,
			StartOfExploitation: event.TimeOfEvent,
		}
		table = club.Tables[event.TableID]
	} else {
		table.Exploitation += table.EndOfExploitation.Sub(table.StartOfExploitation)
		table.StartOfExploitation = event.TimeOfEvent
	}

	//club.Tables[event.TableID] = client.Table
	client.Table = &table
	table.Client = &client
	club.Client[client.ClientID] = client
	club.Tables[table.TableID] = table

}
