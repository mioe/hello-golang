.PHONY: build
build:
		go build -v ./cmd/apiserve

.PHONY: test
test:
		go test -v ./...

.DEFAULT_GOAL := build
