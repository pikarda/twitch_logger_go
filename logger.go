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
		streamer := styledText(UserList[message.Channel][1], message.Channel)
		chatter := styledUser(message.User.Name)
		text := fmt.Sprint(streamer, " ", chatter, ": ", message.Message)
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

	client.OnConnect(func() {
		printChannels()
		go shutdown(client)
	})

	go loop(client)

	err := client.Connect()
	if err != nil {
		fmt.Println(styledStartApp("APP TERMINATED"))
	}

}

func printChannels() {
	var stringUsers string
	for i := range UserList {
		streamer := styledText(UserList[i][1], i)
		stringUsers = fmt.Sprintf("%s%s%s", stringUsers, streamer, ", ")
	}
	fmt.Println("LOGGING CHANNELS: " + stringUsers + "\n\n")
}
