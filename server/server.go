package main

import (
	"github.com/byterotom/octodrive/server/tcp"
	"github.com/byterotom/octodrive/server/udp"
)

func main() {
	go udp.HandleDiscover()
	tcp.RunServer()
}
