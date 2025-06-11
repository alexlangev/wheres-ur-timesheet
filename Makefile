.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o wut ./cmd/tui/main.go

clean:
	go clean
	rm ./wut

test:
	go test -v ./...
