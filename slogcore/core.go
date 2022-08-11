package slogcore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
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
	var callerPayload block

	_, filename, line, ok := runtime.Caller(3)
	if ok {
		callerPayload = block{
			Type: section,
			Fields: []interface{}{
				block{
					Type: markdown,
					Text: fmt.Sprintf("*File*\n%s", filename),
				},
				block{
					Type: markdown,
					Text: fmt.Sprintf("*Line*\n%d", line),
				},
			},
		}
	}

	tagsPayload := block{
		Type: section,
		Text: block{
			Type: markdown,
			Text: strings.Join(tags, " "),
		},
	}

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
			callerPayload,
			tagsPayload,
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
