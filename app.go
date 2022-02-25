package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	//homeDir, _ := os.UserHomeDir()
	exeName := ""
	//
	if goruntime.GOOS == "windows" {
		exeName = "lethean-server.exe"
	} else {
		exeName = "lethean-server"
	}

	//exePath := filepath.Join(homeDir, "Lethean", exeName)

	if _, err := os.Stat(exeName); err == nil {
		fmt.Println("LTHN Found, Starting backend service")
		fmt.Println(exeName)

		if goruntime.GOOS == "windows" {
			cmd := exec.Command(`lethean-server.exe`)
			if _, err := cmd.Output(); err != nil {
				log.Println("Error:", err)
			}
		} else {
			spawnCmd = exec.Command(exeName)
			_, err := spawnCmd.Output()

			if err != nil {
				fmt.Println(err)
				// panic(err)
			}
		}

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Please make sure lethean-server in installed")
		//_ = os.WriteFile(exePath, data, 0777)
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
