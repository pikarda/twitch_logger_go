package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gempir/go-twitch-irc/v4"
)

var disconnect = make(chan os.Signal, 1)

func shutdown(client *twitch.Client) {
	signal.Notify(disconnect, os.Interrupt, syscall.SIGTERM)
	for range disconnect {
		client.Disconnect()
	}
}
