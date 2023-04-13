package slog

import (
	"fmt"
	"log"
	"os"

	"github.com/hakkasuru/slog/core"
)

// Logger struct
type Logger struct {
	config Config
	core   core.Core
}

// NewNoopLogger creates no op logger
func NewNoopLogger() *Logger {
	return &Logger{
		config: Config{},
		core:   core.NewNoopCore(),
	}
}

// NewSlackLogger creates slack logger
func NewSlackLogger(cfg Config) *Logger {
	return &Logger{
		config: cfg,
		core:   core.NewSlackCore(),
	}
}

// Info publish log to slack with info tag
func (l *Logger) Info(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[INFO] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, tags)
}

// Error publish log to slack with error tag
func (l *Logger) Error(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[ERROR] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, tags)
}

// Emergency publish log to slack with emergency tag
func (l *Logger) Emergency(msg string, tags ...string) {
	fmtTitle := fmt.Sprintf("[EMERGENCY] %s", l.config.DefaultTitle)
	l.write(fmtTitle, msg, append(tags, "<!channel>"))
}

func (l *Logger) write(title string, msg string, tags []string) {
	allTags := append(l.config.DefaultTags, tags...)
	err := l.core.Write(l.config.WebhookURL, title, msg, allTags)
	if err != nil {
		log.New(os.Stderr, "[ERROR] ", log.LstdFlags).Printf("%v", err)
	}
}
