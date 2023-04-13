package slog

// Config struct for slack configs
type Config struct {
	WebhookURL   string
	DefaultTitle string
	DefaultTags  []string
}

// NewConfig configuration for slack
func NewConfig(webhookURL string, defaultTitle string, defaultTags ...string) Config {
	return Config{
		WebhookURL:   webhookURL,
		DefaultTitle: defaultTitle,
		DefaultTags:  defaultTags,
	}
}
