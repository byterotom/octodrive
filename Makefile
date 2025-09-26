.PHONY: server
server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./server/
	cd server && docker build -t server:latest .