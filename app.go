package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	debme "github.com/leaanthony/debme"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"syscall"
)

//go:embed server/build
var lthn embed.FS

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

	if _, err := os.Stat(exePath); err == nil {
		fmt.Println("Found Lethean Server: " + exePath)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Could not find Lethean Server, extracting to: " + exePath)
		root, _ := debme.FS(lthn, "server/build")
		err := root.CopyFile(exeName, exePath, 0777)
		if err != nil {
			return
		}
		root = debme.Debme{}
		err = nil
		//_ = os.WriteFile(exePath, data, 0777)
	}

	if goruntime.GOOS == "windows" {
		if _, err := os.Stat(exePath); err == nil {
			fmt.Println("Starting lthn.exe: " + exePath)
			spawnCmd = exec.Command("cmd.exe", "/c", "START", "<Lethean Server> /b /min", exePath, "server")
		} else {
			fmt.Println("Error Could not find lthn.exe:" + exePath)
		}
	} else {
		spawnCmd = exec.Command(exePath, "server")
	}
	cerr := os.Chdir(homeDir)
	if cerr != nil {
		fmt.Println(cerr)
		return
	}
	fmt.Println("Running command and waiting for it to finish...")
	err := spawnCmd.Start()
	if err != nil {
		fmt.Println("cmd.Start failed: ", err)
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
	err := spawnCmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		return
	}
}
