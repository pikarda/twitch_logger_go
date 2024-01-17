package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type twitchData struct {
	Data []struct {
		ID              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageURL string    `json:"profile_image_url"`
		OfflineImageURL string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"data"`
}

// getting an image of the user's avatar with the specified ID
func GetAvatar(id string) (string, error) {
	var link string
	var data *twitchData

	url := fmt.Sprint("https://api.twitch.tv/helix/users?id=", id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", Token))
	req.Header.Add("Client-Id", ClientId)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if err != nil {
			return "", errors.New("error doing request")
		}
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return "", errors.New("GetAvatar ERROR")
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", errors.New("error unmarshaling")
	}

	link = data.Data[0].ProfileImageURL

	return link, nil
}
