package processors

import (
	"fmt"
	"inputfileprocess/helpers"
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

			club.Client[nextClient.ClientID] = model.Client{
				ClientID:    nextClient.ClientID,
				ArrivalTime: nextClient.ArrivalTime,
				Table:       client.Table,
			}

			nextClient = club.Client[nextClient.ClientID]

			nextClient.Table.Client = &nextClient

			client.DepartureTime = event.TimeOfEvent

			client.Table.EndOfExploitation = event.TimeOfEvent
			client.Table.Exploitation += client.Table.StartOfExploitation.Sub(client.Table.EndOfExploitation)
			nextClient.Table.StartOfExploitation = event.TimeOfEvent
			club.Tables[client.Table.TableID] = *client.Table
			fmt.Println(club.Tables[client.Table.TableID].StartOfExploitation, "exploitation", client.ClientID)
			club.HistoryList = append(club.HistoryList, client)
			delete(club.Client, client.ClientID)
			line = fmt.Sprintf("%s %d %s %d\n", event.TimeOfEvent.Format("15:04"), 12, nextClient.ClientID, nextClient.Table.TableID)
			_, err = file.WriteString(line)
			if err != nil {
				fmt.Println("Failed to write to file:", err)
				os.Exit(1)
			}

		} else {
			fmt.Println(client.Table.Exploitation.Seconds(), "exploitation", client.ClientID)
			client.Table.EndOfExploitation = event.TimeOfEvent
			client.Table.Exploitation += client.Table.StartOfExploitation.Sub(client.Table.EndOfExploitation)
			client.DepartureTime = event.TimeOfEvent
			fmt.Println(helpers.GetHours(client.Table.Exploitation), client.ClientID)
			club.HistoryList = append(club.HistoryList, client)
			club.Tables[client.Table.TableID] = *client.Table
			delete(club.Client, client.ClientID)
		}

	}
}
