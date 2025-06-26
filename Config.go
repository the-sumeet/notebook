package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	OnlyMarkdownFiles bool   `json:"onlyMarkdownFiles"`
	NotesDir          string `json:"notesDir"`
}

var config Config

func init() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configPath := homeDir + "/.config/notebook/conf.json"

	config = Config{
		OnlyMarkdownFiles: true,
		NotesDir:          homeDir,
	}

	// Load config from file
	raw, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		panic(err)
	}
}
