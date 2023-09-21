package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {

	client := http.Client{}
	getUpdates := "https://api.telegram.org/bot6194597538:AAGHp55C3b5Alojxh6MYlD6Fix296DWNBQY/getUpdates?offset=%d&timeout=60"
	urlMessage := "https://api.telegram.org/bot6194597538:AAGHp55C3b5Alojxh6MYlD6Fix296DWNBQY/sendMessage"

	id := 1
	for {

		resp, err := client.Get(fmt.Sprintf(getUpdates, id))

		if err != nil {
			fmt.Printf("hello kitty murio :( %s", err)
			return
		}
		decoder := json.NewDecoder(resp.Body)
		response := Response{}
		err = decoder.Decode(&response)
		if err != nil {
			fmt.Printf("hello kitty murio :( %s", err)
			return
		}
		for _, update := range response.Updates {
			fmt.Printf(" %s", update.Message.Text)
			http.PostForm(urlMessage, url.Values{
				"chat_id": []string{fmt.Sprintf("%d", update.Message.Chat.ChatId)},
				"text":    []string{"fyytytyutytyu"},
			})

			id = update.UpdateId + 1
		}
		time.Sleep(1 * time.Second)
	}

}

type Response struct {
	Ok      bool     `json:"ok,omitempty"`
	Updates []Update `json:"result,omitempty"`
}

type Update struct {
	UpdateId int     `json:"update_id,omitempty"`
	Message  Message `json:"message,omitempty"`
}

type Message struct {
	MessageId int    `json:"message_id,omitempty"`
	Text      string `json:"text,omitempty"`
	Chat      Chat   `json:"chat,omitempty"`
}
type Chat struct {
	ChatId int `json:"id,omitempty"`
}
