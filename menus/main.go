package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	// Create application with options
	app := NewApp()

	err := wails.Run(&options.App{
		Title:      "Menus Demo",
		Width:      800,
		Height:     600,
		MinWidth:   400,
		MinHeight:  400,
		MaxWidth:   1280,
		MaxHeight:  1024,
		RGBA:       &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		Menu:       app.applicationMenu(),
		Assets:     assets,
		LogLevel:   logger.DEBUG,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		OnDomReady: app.domready,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
