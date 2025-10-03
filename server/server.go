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
	// Opening log file for writing 
	logFile, err := pkg.OpenFile("logs/server.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// Multiwriter to get logs on terminal and logfile both
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))

	// Changing the default logger
	slog.SetDefault(logger)

	// Starting the discover handler
	go udp.HandleDiscover()
	// Starting the TCP server for file and reqest handling
	tcp.RunServer()
}
