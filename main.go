package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func applicationMenu(deleteNote func(), newNote func(), newFolder func()) *menu.Menu {
	appMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.AppMenu()) // On macOS platform, this must be done right after `NewMenu()`
	}

	FileMenu := appMenu.AddSubmenu("File")
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {

	})

	FileMenu.AddText("Delete", keys.CmdOrCtrl("d"), func(_ *menu.CallbackData) {
		deleteNote()
	})

	FileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		newFolder()
	})

	FileMenu.AddText("New note", keys.CmdOrCtrl("n"), func(_ *menu.CallbackData) {
		newNote()
	})

	FileMenu.AddText("New folder", keys.Combo("n", keys.CmdOrCtrlKey, keys.ShiftKey), func(_ *menu.CallbackData) {
		newFolder()
	})

	return appMenu
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "notebook",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Menu: applicationMenu(app.deleteNote, app.newNote, app.newFolder),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
