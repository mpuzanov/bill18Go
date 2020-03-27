.DEFAULT_GOAL = build
SOURCE=./cmd/bill18go
APP=bill18go
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PKGS := . ./internal
GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

build:
	go build -v -o ${APP} ${SOURCE}

run: 
	go run ${SOURCE} --config=configs/config-prod.yaml

test:
	go test -v -timeout 30s ${GO_TEST_DIRS}

lint:
	@goimports -w ${GO_SRC_DIRS}
	@golangci-lint run

image:
	docker build -t puzanovma/bill18go -f deployments/. 

rundock:
	docker run --rm -it -e "DB_PASSWORD=dnypr1" -p 8090:8090 puzanovma/bill18go

release:
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${APP}.exe ./cmd/${APP}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${APP} ./cmd/${APP}


.DEFAULT_GOAL := build

.PHONY: build run test release lint image rundock
