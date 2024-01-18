package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gempir/go-twitch-irc/v4"
)

var disconnect = make(chan os.Signal, 1)

func shutdown(client *twitch.Client) {
	signal.Notify(disconnect, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGTERM)
	for {
		select {
		case <-disconnect:
			client.Disconnect()
		}
	}
}
