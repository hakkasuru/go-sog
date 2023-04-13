package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type slackCore struct {
	httpClient *http.Client
}

// NewSlackCore slack core implementation
func NewSlackCore() Core {
	return &slackCore{
		httpClient: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
}

// Write message into payload and publish to slack webhook
func (c *slackCore) Write(url string, title string, msg string, tags []string) error {
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

	var blocks blocks
	blocks = append(
		blocks,
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
	)

	if len(tags) > 0 {
		blocks = append(
			blocks,
			block{
				Type: section,
				Text: block{
					Type: markdown,
					Text: strings.Join(tags, " "),
				},
			},
		)
	}

	payload := payload{Blocks: blocks}

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
