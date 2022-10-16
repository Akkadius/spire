package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type WebhookRequestBody struct {
	Content string `json:"content"`
}

func SendDiscordWebhook(webhookUrl string, msg string) error {
	body, _ := json.Marshal(WebhookRequestBody{Content: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
