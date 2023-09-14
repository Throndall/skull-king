package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	client := http.Client{}

	resp, err := client.Get("https://api.telegram.org/bot<TOKEN>/getUpdates")

	if err != nil {
		fmt.Printf("hello kitty murio :( %s", err)
		return
	}

	//	resp.Body.Read()
	decoder := json.NewDecoder(resp.Body)

	response := Response{}

	err = decoder.Decode(&response)
	if err != nil {
		fmt.Printf("hello kitty murio :( %s", err)
		return
	}

	fmt.Printf(" %s", response.Result[0].Message.Text)

}

type Response struct {
	Ok     bool     `json:"ok,omitempty"`
	Result []Update `json:"result,omitempty"`
}

type Update struct {
	UpdateId int     `json:"update_id,omitempty"`
	Message  Message `json:"message,omitempty"`
}

type Message struct {
	MessageId int    `json:"message_id,omitempty"`
	Text      string `json:"text,omitempty"`
}
