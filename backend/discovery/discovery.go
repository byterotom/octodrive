package discovery

import (
	"context"
	"log/slog"
	"net"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Discover(ctx context.Context) {
	laddr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}

	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	conn.SetReadDeadline(time.Now().Add(20 * time.Second))

	raddr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 9999,
	}
	message := "POLLOS"
	conn.WriteToUDP([]byte(message), raddr)

	for {
		buf := make([]byte, 16)
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		if string(buf[:n]) == "OK" {
			slog.Info(src.String())
			runtime.EventsEmit(ctx, "backend:discover", src.String())
		}
	}
}
