package slog

type Config struct {
	WebhookURL   string
	DefaultTitle string
	DefaultTags  []string
}

func NewConfig(webhookURL string, defaultTitle string, defaultTags ...string) Config {
	return Config{
		WebhookURL:   webhookURL,
		DefaultTitle: defaultTitle,
		DefaultTags:  defaultTags,
	}
}
