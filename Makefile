APP=bill18go

build:
	go build -v ./

run: 
	go run .

test:
	go test -v -race -timeout 30s ./...

fmt:
	gofmt -w .

release:
	GOOS=windows GOARCH=amd64 go build -o ${APP}.exe main.go
	GOOS=linux GOARCH=amd64 go build -o ${APP} main.go

.DEFAULT_GOAL := build

.PHONY: build run test release
