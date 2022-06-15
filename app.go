package main

import (
	"bufio"
	"context"
	"embed"
	"errors"
	"fmt"
	debme "github.com/leaanthony/debme"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
)

//go:embed build/cli/deno/bin
var deno embed.FS

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
		exeName = "deno.exe"
	} else {
		exeName = "deno"
	}
	exePath := filepath.Join(cliDir, exeName)

	if _, err := os.Stat(exePath); err == nil {
		fmt.Println("Found Lethean Server: " + exePath)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Could not find Lethean Server, extracting to: " + exePath)
		root, _ := debme.FS(deno, "build/cli/deno/bin")
		err := root.CopyFile(exeName, exePath, 0777)
		if err != nil {
			return
		}
		root = debme.Debme{}
		err = nil
		//_ = os.WriteFile(exePath, data, 0777)
	}

	args := []string{"run", "--unstable", "-A", "https://raw.githubusercontent.com/dAppServer/server/v2/mod.ts"}

	//if goruntime.GOOS == "windows" {
	//	if _, err := os.Stat(exePath); err == nil {
	//		fmt.Println("Starting lthn.exe: " + exePath)
	//		spawnCmd = exec.Command("cmd.exe", "/c", "START", "<Lethean Server> /b /min", exePath, "server")
	//	} else {
	//		fmt.Println("Error Could not find lthn.exe:" + exePath)
	//	}
	//} else {
	spawnCmd = exec.Command(exePath, args...)

	//}
	cerr := os.Chdir(homeDir)
	if cerr != nil {
		fmt.Println(cerr)
		return
	}
	stdout, err := spawnCmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	stderr, err := spawnCmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := spawnCmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		println(scanner.Text())

	}

	_, err = ioutil.ReadAll(stderr)
	if err != nil {
		println(err.Error())
		return
	}
	//if err := spawnCmd.Wait(); err != nil {
	//	println(err.Error())
	//}

	fmt.Println("Running command and waiting for it to finish...")

	//if err != nil {
	//	fmt.Println("cmd.Start failed: ", err)
	//}
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
