.PHONY: all clean build test

all: build

build:
	go build -o hyggemedia .

clean:
	rm -f hyggemedia

test:
	go test ./...