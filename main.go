package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	"path/filepath"
)

//go:embed frontend
var assets embed.FS

var homeDir string

var log logger.Logger

func main() {
	exe, _ := os.Executable()

	if filepath.Base(filepath.Dir(exe)) == "MacOS" {
		homeDir = filepath.Join(filepath.Dir(exe), "../../..")
	} else {
		homeDir = filepath.Dir(exe)
	}
	log = logger.NewFileLogger(filepath.Join(homeDir, "data", "logs", "desktop.log"))
	log.Info(filepath.Dir(exe))
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Lethean",
		Width:  1280,
		Height: 740,
		// MinWidth:          720,
		// MinHeight:         570,
		// MaxWidth:          1280,
		// MaxHeight:         740,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			WebviewUserDataPath:  "Data\\webview",
		},
	})

	if err != nil {
		log.Debug(err.Error())
	}
}
