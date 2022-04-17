package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
)

var spawnCmd *exec.Cmd

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {

	// Perform your setup here
	b.ctx = ctx

	var exeName string
	//
	if goruntime.GOOS == "windows" {
		exeName = "lthn.exe"
	} else {
		exeName = "lthn"
	}
	exePath := filepath.Join(homeDir, exeName)

	if goruntime.GOOS == "windows" {
		if _, err := os.Stat(exePath); err == nil {
			log.Info("Starting lthn.exe: " + exePath)
			spawnCmd = exec.Command("cmd.exe", "/C", "start", "/b", exePath, "server")
		} else {
			log.Debug("Error Could not find lthn.exe:" + exePath)
		}
	} else {
		if _, err := os.Stat(exePath); err == nil {
			log.Info("Starting lthn.exe: " + exePath)
			spawnCmd = exec.Command(exePath, "server")
		}

		if err := spawnCmd.Start(); err != nil {
			log.Debug("lthn.exe error: " + err.Error())
		}
	}

}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
	log.Debug("domReady finished")
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	log.Debug("Desktop Shutdown")
	err := spawnCmd.Process.Kill()
	if err != nil {
		return
	}
}
