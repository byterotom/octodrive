package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"log/slog"
	"net"
)

const CHALLENGE_SIZE = 16

func main() {
	go handleUDP()
	handleTCP()
}

func handleTCP() {
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
		slog.Info("Connection from " + conn.RemoteAddr().String())
		go handshake(conn)
	}
}

func handshake(conn net.Conn) {
	pub := make([]byte, ed25519.PublicKeySize)
	_, err := conn.Read(pub)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	challenge := randomChallenge()
	_, err = conn.Write(challenge)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	signature := make([]byte, ed25519.SignatureSize)
	_, err = conn.Read(signature)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if ed25519.Verify(pub, challenge, signature) {
		slog.Info("Verified....")
	} else {
		slog.Info("Invalid....")
		conn.Close()
	}

}


func handleUDP() {
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 9999,
	}
	
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	
	// slog.Info("Listening on " + conn.LocalAddr().String())
	for {
		buf := make([]byte, 64)
		_, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		slog.Info(string(buf) + " from " + src.String())
		_, err = conn.WriteToUDP([]byte("OK"), src)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func randomChallenge() []byte {
	buf := make([]byte, CHALLENGE_SIZE)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}