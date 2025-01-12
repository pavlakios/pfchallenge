.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: gorun
gorun:
	go run cmd/server/main.go

.PHONY: clean
clean:
	go clean
	rm -f bin/*

.PHONY: test
test:
	go test ./...

.PHONY: benchmark
benchmark:
	go test -bench=. ./...

.PHONY: ws/goapp
ws/goapp:
	curl --include --no-buffer \
	 --header "Connection: Upgrade" \
	 --header "Upgrade: websocket" \
	 --header "Host: localhost:8080" \
	 --header "Origin: http://localhost:8080" \
	 --header "Sec-WebSocket-Key: x3JJHMbDL1EzLkh9GBhXDw==" \
	 --header "Sec-WebSocket-Version: 13" \
	 http://localhost:8080/goapp/ws

.PHONY: cli/sessions
cli/sessions:
	make goapp
	bin/cli sessions -n 5