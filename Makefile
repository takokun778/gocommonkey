.PHONY: run
run:
	@go run main.go

.PHONY: tool
tool:
	@aqua install

.PHONY: lint
lint:
	@golangci-lint run --fix
