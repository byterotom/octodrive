package main

import (
	"embed"
	"os"

	"github.com/byterotom/octodrive/backend"
	"github.com/byterotom/octodrive/pkg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	logFile, err := pkg.OpenFile("logs/octodrive.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	app := backend.NewApp(os.Stdout, logFile)

	err = wails.Run(&options.App{
		Title: "octodrive",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		panic(err)
	}
}
