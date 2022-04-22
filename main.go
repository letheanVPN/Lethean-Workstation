package main

import (
	"embed"
	"fmt"
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
	homeDir, _ = os.UserHomeDir()
	homeDir = filepath.Join(homeDir, "Lethean")

	fmt.Println("HomeDir" + homeDir)
	if _, err := os.Stat(filepath.Join(homeDir, "data", "logs")); err != nil {

		err = os.MkdirAll(filepath.Join(homeDir, "data", "logs"), os.ModePerm)

	}

	log = logger.NewFileLogger(filepath.Join(homeDir, "data", "logs", "desktop.log"))
	fmt.Println("Using Base Directory: " + homeDir)
	// Create an instance of the app structure
	app := NewApp()

	var logLvl logger.LogLevel
	envLog, ok := os.LookupEnv("LTHN_LOG_LEVEL")
	if !ok {
		logLvl = 4
	} else {
		log.Debug("LogLevel adjusted to " + envLog)
		logLvl, _ = logger.StringToLogLevel(envLog)
	}

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
		LogLevel:          logLvl,
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
		fmt.Println("error: ", err)
	}
}
