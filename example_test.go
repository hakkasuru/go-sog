package slog_test

import (
	"testing"

	"github.com/hakkasuru/slog"
)

func TestExampleLogInfo(t *testing.T) {
	config := slog.NewConfig(
		"https://hooks.slack.com/services/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",
		"Test Title",
		"<!here>",
	)

	logger := slog.NewLogger(config)

	logger.Info("test message")
}
