.PHONY: all clean build

all: build

build:
	go build -o hyggemedia ./cmd/hyggemedia

clean:
	rm -f hyggemedia