install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	@echo
	@echo 'Installed tools:'
	@golangci-lint --version