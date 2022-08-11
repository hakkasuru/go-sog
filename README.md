# Slog
Simple Go Slack Logger

# Installation
```shell
go get -u github.com/hakkasuru/slog
```

# Quick Start
```go
config := slog.NewConfig(
    "<webhook url>",
    t.Name(),
    "<!here>",
)
logger := slog.NewSlackLogger(config)
logger.Error("API returned http status %d", 500)
```