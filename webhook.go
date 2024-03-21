package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var myNickname []string

func init() {
	envs := getEnv()

	if val, ok := envs["MY_NICKNAME"]; ok {
		myNickname = strings.Split(val, " ")
	}
}

// post message by using discord webhook
func HookCall(channel string, pic string, message string, username string) error {
	url := UserList[channel]

	if len(myNickname) > 0 {
		for _, v := range myNickname {
			if strings.Contains(strings.ToLower(message), strings.ToLower(v)) {
				message = fmt.Sprintf("@everyone" + message)
			}
		}

	}

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
