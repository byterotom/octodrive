package tcp

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/byterotom/octodrive/pkg"
	"github.com/byterotom/octodrive/server/tcp/requests"
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

	if err := handshake(conn); err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info(fmt.Sprintf("Handshake successful for %s", ip))
	req, err := requests.NewRequest(conn)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	req.HandleConn(conn)
}

func handshake(conn net.Conn) error {
	// Read public key
	pub := make([]byte, ed25519.PublicKeySize)
	_, err := conn.Read(pub)
	if err != nil {
		return err
	}

	// Send a random challenge
	challenge := pkg.RandomChallenge()
	_, err = conn.Write(challenge)
	if err != nil {
		return err
	}

	// Read the signature
	signature := make([]byte, ed25519.SignatureSize)
	_, err = conn.Read(signature)
	if err != nil {
		return err
	}

	// Verify the signature
	if ed25519.Verify(pub, challenge, signature) {
		_, err = conn.Write([]byte("OK"))
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Handshake verification failed")
}
