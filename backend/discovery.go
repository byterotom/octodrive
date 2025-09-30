package backend

import (
	"crypto/ed25519"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) DiscoverDevices() {
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
			if ne, ok := err.(net.Error); !ok || !ne.Timeout() {
				slog.Error(err.Error())
			}
			return
		}
		if string(buf[:n]) == "OK" {
			ip := strings.Split(src.String(), ":")[0]
			runtime.EventsEmit(a.ctx, "backend:discover", ip)
			a.ips = append(a.ips, ip)
		}
	}
}

func (a *App) handshake(rip string) (net.Conn, error) {
	raddr := &net.TCPAddr{
		IP:   net.ParseIP(rip),
		Port: 6969,
	}
	conn, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		return nil, err
	}

	// Send public key
	_, err = conn.Write(a.PublicKey)
	if err != nil {
		return nil, err
	}

	// Receive challenge
	challenge := make([]byte, 1024)
	n, err := conn.Read(challenge)
	if err != nil {
		return nil, err
	}

	// Sign and send the challenge
	signature := ed25519.Sign(a.PrivateKey, challenge[:n])
	_, err = conn.Write(signature)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
