.DEFAULT_GOAL = build
APP=bill18go
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

PKGS := . ./internal

build:
	go build -v ./

run: 
	go run .

test:
	go test -v -race -timeout 30s ./...

lint:
	goimports -local "github.com/mpuzanov/bill18Go" -l . | xargs goimports  -local "github.com/mpuzanov/bill18Go" -w

	goimports -w main.go
	golangci-lint run

image:
	docker build -t puzanovma/bill18go . 

rundock:
	#docker run --rm -it -e "PORT=5000" -v $$(pwd)/logs:/app/logs:rw -p 5000:5000 puzanovma/bill18go
	docker run --rm -it -e "PORT=5000" -v $$(pwd)/logs:/app/logs:rw -p 5000:5000 puzanovma/bill18go-scratch

release:
	#GOOS=windows GOARCH=amd64 go build -o ${APP}.exe main.go
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP} main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${APP} main.go

buildwin:
	go build -ldflags="-H windowsgui"

.DEFAULT_GOAL := build

.PHONY: build run test release
