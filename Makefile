APP_NAME := go-base-books-api
CMD_DIR := cmd/server
BINARY := $(APP_NAME)
PKG := ./...

.PHONY: all build run test lint fmt clean tidy

all: clean build

build:
	go build -o bin/$(BINARY) $(CMD_DIR)/main.go

run:
	go run $(CMD_DIR)/main.go

test:
	go test -v $(PKG)

test-cover:
	go test -cover $(PKG)

fmt:
	go fmt $(PKG)

lint:
	go vet $(PKG)

clean:
	rm -rf bin/

tidy:
	go mod tidy
