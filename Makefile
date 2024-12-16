run: build
	@./bin/gomq


build:
	@go build -o bin/gomq

test:
	go test ./... -v