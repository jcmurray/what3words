all:
	go build ./...
	go test ./...
	go vet ./...

.PHONY: all
