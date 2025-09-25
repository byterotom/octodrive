package backend

import (
	"context"

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
