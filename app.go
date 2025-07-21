package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DirEntry struct {
	Path    string
	Name    string
	IsDir   bool
	Content string
}

type ReturnValue struct {
	DirEntries   []DirEntry
	SelectedNote DirEntry
	Bool         bool
	Error        string
}

// App struct
type App struct {
	config       Config
	ctx          context.Context
	SelectedNote DirEntry
	CurrentDir   string
}

// NewApp creates a new App application struct
func NewApp() *App {

	return &App{
		config:     config,
		CurrentDir: config.NotesDir,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) deleteNote() {
	runtime.EventsEmit(a.ctx, "deleteNote")
}

func (a *App) newNote() {
	runtime.EventsEmit(a.ctx, "newNote")
}

func (a *App) newFolder() { runtime.EventsEmit(a.ctx, "newFolder") }

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SetTitle(ctx context.Context, title string) {
	runtime.WindowSetTitle(ctx, title)
}

func (a *App) GetSelectedNote() DirEntry {
	return a.SelectedNote
}

func (a *App) SelectNote(noteFileName string) ReturnValue {
	notePath := filepath.Join(a.CurrentDir, noteFileName)
	fmt.Println(notePath)

	file, err := os.OpenFile(notePath, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return ReturnValue{
				Error: fmt.Sprintf("File %s does not exist", notePath),
			}
		}
		panic(err)
	}
	defer file.Close()

	// Check if the path is a directory
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		return ReturnValue{
			Error: fmt.Sprintf("Path %s is a directory, not a file", notePath),
		}
	}

	a.SelectedNote = DirEntry{
		Path:    notePath,
		Name:    info.Name(),
		IsDir:   false,
		Content: "",
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return ReturnValue{
		SelectedNote: DirEntry{
			Path:    notePath,
			Name:    info.Name(),
			IsDir:   false,
			Content: string(content),
		},
	}
}

func (a *App) SaveFileContent(fileName string, content string) ReturnValue {
	filePath := filepath.Join(a.CurrentDir, fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return ReturnValue{
			Error: fmt.Sprintf("Error opening file %s: %v", filePath, err),
		}
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return ReturnValue{
			Error: fmt.Sprintf("Error writing to file %s: %v", filePath, err),
		}
	}

	return ReturnValue{}
}

func (a *App) GoBack() ReturnValue {
	// Get the parent directory
	parentDir := filepath.Dir(a.CurrentDir)

	// Check if the parent directory is the same as the current directory
	if parentDir == a.CurrentDir {
		return ReturnValue{
			Error: "Already at the root directory",
		}
	}

	a.CurrentDir = parentDir

	return ReturnValue{
		SelectedNote: DirEntry{
			Path:    filepath.Join(a.CurrentDir, a.SelectedNote.Name),
			Name:    a.SelectedNote.Name,
			IsDir:   false,
			Content: "",
		},
	}
}

func (a *App) SelectDir(dirName string) ReturnValue {
	dirPath := filepath.Join(a.CurrentDir, dirName)

	file, err := os.OpenFile(dirPath, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return ReturnValue{
				Error: fmt.Sprintf("Directory %s does not exist", dirPath),
			}
		}
		panic(err)
	}
	defer file.Close()
	// Check if the path is a directory
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		return ReturnValue{
			Error: fmt.Sprintf("Path %s is not a directory", dirPath),
		}
	}
	a.CurrentDir = dirPath
	runtime.WindowSetTitle(a.ctx, dirPath)

	return ReturnValue{}
}

func (a *App) GetFiles() ReturnValue {

	dirItems, _ := os.ReadDir(a.CurrentDir)
	dirEntries := []DirEntry{}
	for _, entry := range dirItems {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		dirEntries = append(dirEntries, DirEntry{
			Path:  filepath.Join(a.CurrentDir, entry.Name()),
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}
	return ReturnValue{
		DirEntries: dirEntries,
	}
}

func (a *App) NewNote(name string) ReturnValue {

	notePath := filepath.Join(a.CurrentDir, name)
	file, err := os.OpenFile(notePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return ReturnValue{
			Error: fmt.Sprintf("Error creating file %s: %v", notePath, err),
		}
	}
	defer file.Close()

	return ReturnValue{
		SelectedNote: DirEntry{
			Path:    notePath,
			Name:    name,
			IsDir:   false,
			Content: "",
		},
	}

}

func (a *App) DeleteNoteFile(notePath string) ReturnValue {
	err := os.Remove(notePath)
	if err != nil {
		return ReturnValue{
			Error: fmt.Sprintf("Error deleting file %s: %v", notePath, err),
		}
	}

	return ReturnValue{}
}

func (a *App) NewFolder(name string) ReturnValue {
	folderPath := filepath.Join(a.CurrentDir, name)
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		return ReturnValue{
			Error: fmt.Sprintf("Error creating folder %s: %v", folderPath, err),
		}
	}

	return ReturnValue{}
}

func (a *App) IsNotesDir() ReturnValue {
	fmt.Println("Current Directory:", a.CurrentDir, a.config.NotesDir)
	return ReturnValue{
		Bool: a.CurrentDir == a.config.NotesDir,
	}

}
