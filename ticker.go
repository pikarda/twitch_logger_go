package main

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

var ticker = time.NewTicker(5 * time.Minute)

// function that tracks if app needs to reload list of streamers (if value of reload == 1)
func loop(client *twitch.Client) {
	for {
		select {
		case <-ticker.C:
			GetReloadStatus()
			if ReloadStatus == true {
				for _, v := range listOfChannels {
					client.Depart(v)
					fmt.Printf("Channel departed %s \n", v)
				}
				listOfChannels = []string{}
				UserList = map[string][]string{}
				GetNotionData()
				setVariables(Data)
				client.Join(listOfChannels...)
				printChannels()
			}
		}
	}
}
