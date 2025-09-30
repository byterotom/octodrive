package backend

import (
	"log/slog"
	"net"
)

func (a *App) SendFile(filePath string) {
	conns := make([]net.Conn, len(a.ips))
	for idx, ip := range a.ips {
		var err error
		conns[idx], err = a.handshake(ip)
		if err != nil {
			slog.Error(err.Error())
		}
		conns[idx].Write([]byte{1 << 2})
	}
}
