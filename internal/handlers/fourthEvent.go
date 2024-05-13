package handlers

import (
	"fmt"
	"github.com/aleksandra0KR/input-file-processing-game-club/internal/helpers"
	model2 "github.com/aleksandra0KR/input-file-processing-game-club/internal/models"
	"os"
)

func FourthEvent(event *model2.Event, club *model2.Club, file *os.File) {
	line := fmt.Sprintf("%s %d %s\n", event.TimeOfEvent.Format("15:04"), event.EventID, event.ClientID)
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
	} else if client.Table == nil {
		delete(club.Client, client.ClientID)
	} else {
		if len(club.WaitingList) > 0 {
			nextClient := club.WaitingList[0]
			club.WaitingList = club.WaitingList[1:]

			club.Client[nextClient.ClientID] = model2.Client{
				ClientID:    nextClient.ClientID,
				ArrivalTime: nextClient.ArrivalTime,
				Table:       client.Table,
			}

			nextClient = club.Client[nextClient.ClientID]
			nextClient.Table.Client = &nextClient
			client.DepartureTime = event.TimeOfEvent
			client.Table.EndOfExploitation = event.TimeOfEvent
			client.Table.Exploitation += client.Table.StartOfExploitation.Sub(client.Table.EndOfExploitation)
			client.Table.Payment += helpers.GetHours(client.Table.StartOfExploitation.Sub(client.Table.EndOfExploitation))
			nextClient.Table.StartOfExploitation = event.TimeOfEvent
			club.Tables[client.Table.TableID] = *client.Table
			delete(club.Client, client.ClientID)

			line = fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent.Format("15:04"), 12, nextClient.ClientID, nextClient.Table.TableID)
			_, err = file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}

		} else {
			table := client.Table
			table.EndOfExploitation = event.TimeOfEvent
			table.Exploitation += table.StartOfExploitation.Sub(table.EndOfExploitation)
			table.Payment += helpers.GetHours(table.StartOfExploitation.Sub(table.EndOfExploitation))
			client.DepartureTime = event.TimeOfEvent
			table.Client = nil
			club.Tables[client.Table.TableID] = *table
			delete(club.Client, client.ClientID)
		}

	}
}
