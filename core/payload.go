package core

const (
	header    = "header"
	section   = "section"
	markdown  = "mrkdwn"
	plaintext = "plain_text"
)

type blocks []block

type payload struct {
	Blocks blocks `json:"blocks"`
}

type block struct {
	Type   string        `json:"type"`
	Text   interface{}   `json:"text,omitempty"`
	Fields []interface{} `json:"fields,omitempty"`
}
