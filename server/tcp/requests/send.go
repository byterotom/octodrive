package requests

import (
	"log/slog"
	"net"
)

type Send struct {
}

func (s *Send) HandleConn(conn net.Conn) {
	slog.Info("Send signal received !")
}
