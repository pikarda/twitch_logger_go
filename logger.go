package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

var listOfChannels []string

// func Logger(data *NotionData) {
func Logger() {
	setVariables(Data)
	for i := range UserList {
		listOfChannels = append(listOfChannels, i)
	}
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		//check if user is bot
		if slices.Contains(BlackList, strings.ToLower(message.User.Name)) {
			return
		}

		text := fmt.Sprint("#", message.Channel, "- ", message.User.Name, ": ", message.Message)
		fmt.Println(text)
		pic, err := GetAvatar(message.User.ID)
		if err != nil {
			fmt.Println(err)
		}
		err = HookCall(message.Channel, pic, message.Message, message.User.Name)
		if err != nil {
			fmt.Println(err)
		}
	})

	client.Join(listOfChannels...)

	client.OnConnect(printChannels)

	go loop(client)

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}

func printChannels() {
	var stringUsers string
	for i := range UserList {
		stringUsers = fmt.Sprintf("%s%s%s", stringUsers, i, ", ")
	}
	fmt.Println("LOGGING CHANNELS: " + stringUsers)
}
