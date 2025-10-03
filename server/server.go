package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/byterotom/octodrive/pkg"
	"github.com/byterotom/octodrive/server/tcp"
	"github.com/byterotom/octodrive/server/udp"
)

func main() {

	logFile, err := pkg.OpenFile("logs/server.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))

	slog.SetDefault(logger)

	go udp.HandleDiscover()
	tcp.RunServer()
}
