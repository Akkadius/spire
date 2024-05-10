package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// SendMessage sends a message to a discord webhook
func SendMessage(webhookUrl, header, contents string) {
	chunkCount := 1
	for _, chunk := range chunkString(contents, 1800) {
		_ = SendDiscordWebhook(
			webhookUrl,
			header+fmt.Sprintf(" **Chunk** [%v]\n```\n%v\n```", chunkCount, chunk),
		)
		chunkCount++
	}
}

// chunkSubstr splits a string into chunks of a given size
func chunkString(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}
