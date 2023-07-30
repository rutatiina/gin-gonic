package services

import (
	"bytes"
	"net/http"
)

func LogToSlack(msg string) {
	url := "https://hooks.slack.com/services/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"

	jsonStr := []byte(`{"text":"` + msg + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle response
}
