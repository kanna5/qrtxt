build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o qrtxt ./cmd

.PHONY: build
