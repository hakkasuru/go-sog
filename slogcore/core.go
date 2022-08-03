package slogcore

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Core interface {
	Write(url string, title string, msg string, tags []string) error
}

type noopCore struct{}

func NewNoopCore() Core                                                            { return noopCore{} }
func (c noopCore) Write(url string, title string, msg string, tags []string) error { return nil }

type core struct {
	httpClient *http.Client
}

func NewCore() Core {
	return &core{
		httpClient: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
}

func (c *core) Write(url string, title string, msg string, tags []string) error {
	payload := payload{
		Blocks: blocks{
			block{
				Type: header,
				Text: block{
					Type: plaintext,
					Text: title,
				},
			},
			block{
				Type: section,
				Text: block{
					Type: markdown,
					Text: msg,
				},
			},
			block{
				Type: section,
				Text: block{
					Type: markdown,
					Text: strings.Join(tags, " "),
				},
			},
		},
	}

	data, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return marshalErr
	}

	req, reqErr := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if reqErr != nil {
		return reqErr
	}

	req.Header.Add("Content-Type", "application/json")

	_, resErr := c.httpClient.Do(req)
	if resErr != nil {
		return resErr
	}

	return nil
}
