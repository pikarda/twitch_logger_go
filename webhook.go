package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Channels struct {
	Channel string
	webhook string
}

// post message by using discord webhook
func HookCall(channel string, pic string, message string, username string) error {

	url := UserList[channel]

	payloadString := fmt.Sprint("content=", message, "&username=", username, "&avatar_url=", pic)

	payload := strings.NewReader(payloadString)

	req, _ := http.NewRequest("POST", url[0], payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		return errors.New("HookCall ERROR")
	}

	defer res.Body.Close()

	return nil
}
