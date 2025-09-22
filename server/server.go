package server

import (
	"context"
	"octodrive/server/auth"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const SECRET_PHRASE_SIZE = 12

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
		secretPhrase, err := loadSecretPhrase()
		if err != nil {
			runtime.EventsEmit(ctx, "server:showSetup")
			return
		}

		a.Auth = auth.NewAuth(secretPhrase)
		runtime.EventsEmit(ctx, "server:showHome")
	})
}

func (a *App) GenerateSecretPhrase() string {

	if a.Auth != nil {
		return a.Auth.SecretPhrase
	}

	secretPhrase := auth.NewSecretPhrase()
	auth.SaveSecretPhraseOnSystem(secretPhrase)
	a.Auth = auth.NewAuth(secretPhrase)

	return secretPhrase
}

func (a *App) SetSecretPhrase(secretPhrase string) {
	auth.SaveSecretPhraseOnSystem(secretPhrase)
	a.Auth = auth.NewAuth(secretPhrase)
}

func loadSecretPhrase() (string, error) {
	buf, err := os.ReadFile("server/data/secret_phrase.txt")
	return string(buf), err
}
