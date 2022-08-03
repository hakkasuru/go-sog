package slog_test

import (
	"testing"

	"github.com/hakkasuru/slog"
)

func TestExampleLogInfo(t *testing.T) {
	config := slog.NewConfig(
		"https://hooks.slack.com/services/webhook", // replace with webhook url
		t.Name(),
		"<!here>",
	)

	logger := slog.NewLogger(config)

	logger.Info("test message")
}

func TestExampleErrorInfo(t *testing.T) {
	config := slog.NewConfig(
		"https://hooks.slack.com/services/webhook", // replace with webhook url
		t.Name(),
		"<!here>",
	)

	logger := slog.NewLogger(config)

	logger.Error("test message")
}

func TestExampleEmergencyInfo(t *testing.T) {
	config := slog.NewConfig(
		"https://hooks.slack.com/services/webhook", // replace with webhook url
		t.Name(),
		"<!here>",
	)

	logger := slog.NewLogger(config)

	logger.Emergency("test message")
}
