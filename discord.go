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

func (webhook *Webhook) sentMessage(content string) error {
	reqBody, err := json.Marshal(map[string]any{
		"content": content,
	})
	if err != nil {
		return err
	}

	re, err := http.Post(webhook.Url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(re)
		return err
	}

	return nil
}
