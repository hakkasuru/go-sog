package slog

import (
	"fmt"

	"github.com/hakkasuru/slog/core"
)

type Logger struct {
	config Config
	core   core.Core
}

func NewNoopLogger() *Logger {
	return &Logger{
		config: Config{},
		core:   core.NewNoopCore(),
	}
}

func NewLogger(cfg Config) *Logger {
	return &Logger{
		config: cfg,
		core:   core.NewCore(),
	}
}

func (l *Logger) Info(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[INFO] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, tags)
}

func (l *Logger) Error(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[ERROR] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, tags)
}

func (l *Logger) Emergency(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[EMERGENCY] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, append(tags, "<!channel>"))
}

func (l *Logger) write(title string, msg string, tags []string) {
	allTags := append(l.config.DefaultTags, tags...)
	l.core.Write(l.config.WebhookURL, title, msg, allTags)
}
