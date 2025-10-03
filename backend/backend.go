package backend

import (
	"context"
	"io"
	"log/slog"

	"github.com/byterotom/octodrive/backend/auth"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	*auth.Auth
	ctx context.Context
	ips []string

	logger *slog.Logger
}

func NewApp(logFiles ...io.Writer) *App {

	// Multiwriter to get logs on terminal and logfile both
	multiwriter := io.MultiWriter(logFiles...)

	return &App{
		Auth:   nil,
		ips:    []string{},
		logger: slog.New(slog.NewTextHandler(multiwriter, nil)),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// Setup a event listner to listen when the fromtend is ready
	runtime.EventsOnce(ctx, "frontend:checkSecret", func(...any) {
		secretPhrase, err := auth.LoadSecretPhraseFromSystem()
		if err != nil {
			// Show setup page to generate a new secret phrase if not already on system
			runtime.EventsEmit(ctx, "backend:showSetup")
			return
		}

		a.Auth = auth.NewAuth(secretPhrase)
		runtime.EventsEmit(ctx, "backend:showHome")
	})
}

func (a *App) GenerateSecretPhrase() (string, error) {

	// Return the existing phrase from auth instance
	if a.Auth != nil {
		return a.SecretPhrase, nil
	}

	secretPhrase, err := auth.NewSecretPhrase()
	if err != nil {
		a.logger.Error(err.Error())
		return "", err
	}

	// Setup a new auth instance
	a.Auth = auth.NewAuth(secretPhrase)

	return secretPhrase, nil
}

func (a *App) SetSecretPhrase(secretPhrase string) error {
	// Setup a new auth instance
	a.Auth = auth.NewAuth(secretPhrase)

	if err := auth.SaveSecretPhraseOnSystem(secretPhrase); err != nil {
		a.logger.Error(err.Error())
		return err
	}
	return nil
}
