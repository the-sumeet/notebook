package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DirEntry struct {
	Path  string
	Name  string
	IsDir bool
}

type ReturnValue struct {
	DirEntries []DirEntry
	Error      string
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SetTitle(ctx context.Context, title string) {
	runtime.WindowSetTitle(ctx, title)
}

func (a *App) GetFiles() ReturnValue {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dirItems, _ := os.ReadDir(homeDir)
	dirEntries := []DirEntry{}
	for _, entry := range dirItems {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		dirEntries = append(dirEntries, DirEntry{
			Path:  fmt.Sprintf("%s/%s", homeDir, entry.Name()),
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}
	return ReturnValue{
		DirEntries: dirEntries,
	}
}
