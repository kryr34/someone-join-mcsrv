package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Webhook struct {
	Url string
}

func (webhook *Webhook) sentMessage(content string) {
	reqBody, err := json.Marshal(map[string]any{
		"content": content,
	})
	FatalIfErr(err)

	_, err = http.Post(webhook.Url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
	}
}
