package handlers

import (
	"fmt"
	model2 "inputfileprocess/internal/models"
	"os"
)

func SecondEvent(event *model2.Event, club *model2.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID, event.TableID)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		os.Exit(1)
	}
	club.HistoryList = append(club.HistoryList, event.TimeOfEvent)

	client, ok := club.Client[event.ClientID]
	if !ok {
		line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "ClientUnknown")
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			os.Exit(1)
		}
		return
	}
	table, ok := club.Tables[event.TableID]
	if !ok {
		club.Tables[event.TableID] = model2.Table{
			TableID:             event.TableID,
			Client:              &client,
			StartOfExploitation: event.TimeOfEvent,
		}
		table = club.Tables[event.TableID]
	} else {
		if table.Client != nil {

			line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), 13, "PlaceIsBusy")
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}
		} else {
			table.Exploitation += table.EndOfExploitation.Sub(table.StartOfExploitation)
			table.StartOfExploitation = event.TimeOfEvent
		}

	}

	client.Table = &table
	table.Client = &client
	club.Client[client.ClientID] = client
	club.Tables[table.TableID] = table

}
