package backend

import (
	"context"
	"crypto/ed25519"
	"log/slog"
	"net"

	"github.com/byterotom/octodrive/backend/auth"
	"github.com/byterotom/octodrive/backend/discovery"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const SECRET_PHRASE_SIZE = 10

type App struct {
	*auth.Auth
	ctx context.Context
}

func NewApp() *App {
	return &App{Auth: nil}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	runtime.EventsOnce(ctx, "frontend:checkSecret", func(...any) {
		secretPhrase, err := auth.LoadSecretPhraseFromSystem()
		if err != nil {
			runtime.EventsEmit(ctx, "backend:showSetup")
			return
		}

		a.Auth = auth.NewAuth(secretPhrase)
		runtime.EventsEmit(ctx, "backend:showHome")
	})
}

func (a *App) GenerateSecretPhrase() string {

	if a.Auth != nil {
		return a.SecretPhrase
	}

	secretPhrase := auth.NewSecretPhrase()
	a.Auth = auth.NewAuth(secretPhrase)

	return secretPhrase
}

func (a *App) SetSecretPhrase(secretPhrase string) {
	if a.Auth == nil {
		a.Auth = auth.NewAuth(secretPhrase)
	}
	auth.SaveSecretPhraseOnSystem(secretPhrase)
}

func (a *App) DiscoverDevices() {
	discovery.Discover(a.ctx)
}

func (a *App) ConnectServer(rip string) {
	raddr := &net.TCPAddr{
		IP:   net.ParseIP(rip),
		Port: 6969,
	}
	conn, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer conn.Close()

	_, err = conn.Write(a.PublicKey)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	challenge := make([]byte, 1024)
	n, err := conn.Read(challenge)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	signature := ed25519.Sign(a.PrivateKey, challenge[:n])
	_, err = conn.Write(signature)
	if err != nil {
		slog.Error(err.Error())
		return
	}

}
