package backend

import (
	"context"

	"github.com/byterotom/octodrive/backend/auth"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	*auth.Auth
	ctx context.Context
	ips []string
}

func NewApp() *App {
	return &App{Auth: nil, ips: []string{}}
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

func (a *App) GenerateSecretPhrase() (string, error) {

	if a.Auth != nil {
		return a.SecretPhrase, nil
	}

	secretPhrase, err := auth.NewSecretPhrase()
	if err != nil {
		return "", err
	}

	a.Auth = auth.NewAuth(secretPhrase)

	return secretPhrase, nil
}

func (a *App) SetSecretPhrase(secretPhrase string) error {
	a.Auth = auth.NewAuth(secretPhrase)
	return auth.SaveSecretPhraseOnSystem(secretPhrase)
}
