package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	"path/filepath"
)

//go:embed frontend/dist
var assets embed.FS

var homeDir string
var exeDir string
var cliDir string

var log logger.Logger

func main() {
	homeDir, _ = os.UserHomeDir()
	exeDir, _ = os.Executable()
	homeDir = filepath.Join(homeDir, "Lethean")
	cliDir = filepath.Join(homeDir, "cli")

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
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		LogLevel:          logLvl,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Maximised,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 true,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "Lethean Desktop",
				Message: "© Lethean Community Interest Company",
			},
		},
	})

	if err != nil {
		fmt.Println("error: ", err)
	}
}
