package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/leaanthony/debme"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"syscall"
)

//go:embed server/build
var dappserver embed.FS

var spawnCmd *exec.Cmd

// App struct
type App struct {
	ctx context.Context
}

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
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
		exeName = "server.exe"
	} else {
		exeName = "server"
	}
	exePath := filepath.Join(homeDir, exeName)

	if _, err := os.Stat(exePath); err == nil {
		fmt.Println("Found Lethean Server: " + exePath)
	} else {
		fmt.Println("Could not find Lethean Server, extracting to: " + exePath)
		root, _ := debme.FS(dappserver, "server/build")
		err := root.CopyFile(exeName, exePath, 0777)
		if err != nil {
			return
		}
		root = debme.Debme{}
		err = nil
		//_ = os.WriteFile(exePath, data, 0777)
	}

	cerr := os.Chdir(homeDir)

	if cerr != nil {
		fmt.Println(cerr)
		return
	}
	fmt.Println("Working Directory:" + homeDir)

	if goruntime.GOOS == "windows" {
		if _, err := os.Stat(exePath); err == nil {
			fmt.Println("Starting dappserver.exe: " + exePath)
			spawnCmd = exec.Command("cmd.exe", "/c", "START", "<dAppServer> /b /min", exePath)
		} else {
			fmt.Println("Error Could not find dappserver.exe:" + exePath)
		}
	} else {
		fmt.Println("Starting dappserver: " + exePath)
		spawnCmd = exec.Command(exePath)
	}
	if proc, err := Start(exePath, "start"); err == nil {
		fmt.Println("Started dappserver: ")
		if err != nil {
			fmt.Println("cmd.Start failed: ", proc)
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
	err := spawnCmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		return
	}
}

// GetUserSelectedDirectoryPath returns the path of the directory selected by the user
func (b *App) GetUserSelectedDirectoryPath() string {
	dir, err := runtime.OpenDirectoryDialog(b.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		return homeDir
	}
	fmt.Println("Selected Dir:" + dir)
	return dir
}

// GetUserSelectedFilePath returns the path of the file selected by the user
func (b *App) GetUserSelectedFilePath() string {
	file, err := runtime.OpenFileDialog(b.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		return homeDir + "/wallet"
	}
	fmt.Println("Selected File:" + file)
	return file
}

// GetUserSelectedSaveFilePath returns the path of the file selected by the user
func (b *App) GetUserSelectedSaveFilePath(defaultPath string) string {
	file, err := runtime.SaveFileDialog(b.ctx, runtime.SaveDialogOptions{DefaultDirectory: defaultPath})

	if err != nil {
		return homeDir + "/wallet"
	}
	fmt.Println("Selected File:" + file)
	return file
}
