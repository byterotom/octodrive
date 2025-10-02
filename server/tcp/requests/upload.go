package requests

import (
	"log/slog"
	"net"
)

type Upload struct {
}

func (u *Upload) HandleConn(conn net.Conn) {
	slog.Info("Upload signal received !")
}
