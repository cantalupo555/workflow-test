VERSION ?= dev

.PHONY: build clean

build:
	go build -ldflags="-X main.appVersion=$(VERSION)" -o bin/workflow-test ./cmd/app

clean:
	rm -rf bin/
