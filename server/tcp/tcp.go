package tcp

import (
	"crypto/ed25519"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/byterotom/octodrive/pkg"
)

func RunServer() {
	laddr := &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 6969,
	}
	listner, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	for {
		conn, err := listner.AcceptTCP()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	ip := strings.Split(conn.RemoteAddr().String(), ":")[0]

	if !handshake(conn) {
		slog.Info(fmt.Sprintf("Handshake failed for %s", ip))
		return
	}
	slog.Info(fmt.Sprintf("Handshake successful for %s", ip))
}

func handshake(conn net.Conn) bool {
	// Read public key
	pub := make([]byte, ed25519.PublicKeySize)
	_, err := conn.Read(pub)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	// Send a random challenge
	challenge := pkg.RandomChallenge()
	_, err = conn.Write(challenge)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	// Read the signature
	signature := make([]byte, ed25519.SignatureSize)
	_, err = conn.Read(signature)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	// Verify the signature
	return ed25519.Verify(pub, challenge, signature)
}
