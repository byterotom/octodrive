# Script to build the server binary
.PHONY: server
server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server/bin/server ./server/