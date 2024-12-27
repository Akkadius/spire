package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// SendCrashMessage sends a message to a discord webhook
func SendCrashMessage(webhookUrl, header, contents string) {
	chunkCount := 1
	for _, chunk := range chunkString(contents, 1800) {
		chunk := strings.ReplaceAll(chunk, "`", "\\`")
		_ = SendDiscordWebhook(
			webhookUrl,
			fmt.Sprintf(
				"**%v** **Chunk** [%v]\n```bash\n%v\n```",
				header,
				chunkCount,
				chunk,
			),
		)

		//fmt.Println(header + fmt.Sprintf(" **Chunk** [%v]\n```\n%v\n```", chunkCount, chunk))
		chunkCount++
	}
}

// chunkSubstr splits a string into chunks of a given size
func chunkString(s string, chunkSize int) []string {
	if len(s) == 0 {
		return []string{}
	}

	var chunks []string
	var currentChunk string
	lines := strings.Split(s, "\n") // Split input into lines

	for _, line := range lines {
		// Check if adding this line would exceed the chunk size
		if len(currentChunk)+len(line)+1 > chunkSize { // +1 accounts for the newline
			chunks = append(chunks, currentChunk)
			currentChunk = ""
		}

		// Add the line to the current chunk
		if currentChunk != "" {
			currentChunk += "\n"
		}
		currentChunk += line
	}

	// Add the last chunk if it's not empty
	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
