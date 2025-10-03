package udp

import (
	"log/slog"
	"net"
)

func HandleDiscover() {
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 9999,
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for {
		buf := make([]byte, 64)
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			slog.Error(err.Error())
		}

		slog.Info(string(buf[:n]) + " from " + src.String())
		_, err = conn.WriteToUDP([]byte("OK"), src)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}
