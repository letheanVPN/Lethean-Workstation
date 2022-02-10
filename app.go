package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/leaanthony/debme"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	homeDir, _ := os.UserHomeDir()
	exeName := ""
	//
	if goruntime.GOOS == "windows" {
		exeName = "lethean-server.exe"
	} else {
		exeName = "lethean-server"
	}
	// make folder $HOME/Lethean

	exePath := filepath.Join(homeDir, "Lethean", exeName)

	if _, err := os.Stat(exePath); err == nil {
	    fmt.Println("LTHN Found, Starting backend service")
        spawnCmd = exec.Command(exePath, "backend", "start")
	} else if errors.Is(err, os.ErrNotExist) {
	    fmt.Println("Please make sure lethean-server in installed")
		//_ = os.WriteFile(exePath, data, 0777)
	}
	_, err := spawnCmd.Output()
	if err != nil {
	    fmt.Println(err)
		 panic(err)
	}

}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	err := spawnCmd.Process.Kill()
	if err != nil {
		return
	}
}



// Shows a Dialog
func (b *App) ShowDialog() {
	_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Native Dialog from Go",
		Message: "This is a Native Dialog send from Go.",
	})

	if err != nil {
		panic(err)
	}
}
