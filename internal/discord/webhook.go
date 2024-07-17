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

		//fmt.Println(header + fmt.Sprintf(" **Chunk** [%v]\n```\n%v\n```", chunkCount, chunk))
		chunkCount++
	}
}

// chunkSubstr splits a string into chunks of a given size
func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
