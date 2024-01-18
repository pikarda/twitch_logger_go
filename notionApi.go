package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

//go:embed .env
var env []byte

// getting all secrets from .env file and put in map
func getEnv() map[string]string {
	envs, err := godotenv.Unmarshal(string(env))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return envs
}

type NotionData struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  any `json:"cover"`
		Icon   any `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			Color struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    any    `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      any    `json:"href"`
				} `json:"rich_text"`
			} `json:"color"`
			Value struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    any    `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      any    `json:"href"`
				} `json:"rich_text"`
			} `json:"value"`
			Tag struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"tag"`
			Name struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    any    `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      any    `json:"href"`
				} `json:"title"`
			} `json:"name"`
		} `json:"properties"`
		URL       string `json:"url"`
		PublicURL any    `json:"public_url"`
	} `json:"results"`
	NextCursor     any    `json:"next_cursor"`
	HasMore        bool   `json:"has_more"`
	Type           string `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	RequestID string `json:"request_id"`
}

var Token string
var ClientId string

var UserList = make(map[string][]string)

var BlackList []string

var ReloadStatus bool
var ReloadPageID string

var Data *NotionData

// get data from notion database
func GetNotionData() {
	envs := getEnv()
	var data *NotionData

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprint("https://api.notion.com/v1/databases/"+envs["DB_ID"]+"/query"), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprint("Bearer "+envs["NOTION_TOKEN"]))
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		log.Fatal(err)
	}

	err = UpdateDatabaseReloadStatus()
	if err != nil {
		log.Fatal(err)
	}
	Data = data
	// return data
}

func setVariables(data *NotionData) {
	for _, v := range data.Results {
		if v.Properties.Tag.Select.Name == "secret" {
			if v.Properties.Name.Title[0].Text.Content == "TOKEN" {
				Token = v.Properties.Value.RichText[0].Text.Content
			}
			if v.Properties.Name.Title[0].Text.Content == "CLIENT_ID" {
				ClientId = v.Properties.Value.RichText[0].Text.Content
			}
		}

		if v.Properties.Tag.Select.Name == "user" {
			if len(v.Properties.Color.RichText) > 0 {
				UserList[v.Properties.Name.Title[0].Text.Content] = []string{v.Properties.Value.RichText[0].Text.Content, v.Properties.Color.RichText[0].Text.Content}
			} else {
				UserList[v.Properties.Name.Title[0].Text.Content] = []string{v.Properties.Value.RichText[0].Text.Content, "#AD8E77"}
			}
		}

		if v.Properties.Tag.Select.Name == "blacklist" {
			value := v.Properties.Value.RichText[0].Text.Content
			BlackList = strings.Split(value, ",")
			for i := range BlackList {
				BlackList[i] = strings.TrimSpace(BlackList[i])
			}
		}
	}
	for i := range UserList {
		listOfChannels = append(listOfChannels, i)
	}
}

func GetReloadStatus() {
	envs := getEnv()
	var data *NotionData

	url := fmt.Sprint("https://api.notion.com/v1/databases/" + envs["DB_ID"] + "/query")

	payload := strings.NewReader("{\n  \"filter\": {\n    \"property\": \"tag\",\n    \"select\": {\n      \"equals\": \"reload\"\n    }\n  }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer "+envs["NOTION_TOKEN"]))
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	ReloadPageID = data.Results[0].ID

	if data.Results[0].Properties.Value.RichText[0].Text.Content == "1" {
		ReloadStatus = true
	} else {
		ReloadStatus = false
	}
}

// set reload status to 0 (zero)
func UpdateDatabaseReloadStatus() error {

	url := fmt.Sprint("https://api.notion.com/v1/pages/" + ReloadPageID)

	payload := strings.NewReader("{\n  \"properties\": {\n    \"value\": {\n      \"rich_text\": [\n        {\n          \"type\": \"text\",\n          \"text\": {\n            \"content\": \"0\"\n          }\n        }\n      ]\n    }\n  }\n}")

	req, err := http.NewRequest("PATCH", url, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer secret_KWwleNo3kCCAyFYe5wMWrhIFPmGe4BAvLr40h3XTWdp")
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("Update reload status in DB ERROR")
	}
	ReloadStatus = false
	return nil
}
