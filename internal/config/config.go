package config

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/yousefakbar/o/internal/obsidian"
)

const (
	envObsidianVaultPath	= "OBSIDIAN_VAULT_PATH"
	envEditor		= "EDITOR"
)

type ConfigManager struct {
	// Environment variables related to program behavior
	Editor			string

	// Fields related to Obsidian metadata and settings
	ObsidianVaultPath	string
	DotObsidianVaultPath	string

	// Configuration(s) related to program behavior
	DailyNotesConfig	*obsidian.DailyNotesConfig
	NewFileConfig		*obsidian.NewFileConfig
}

// LoadConfig: loads configuration from environment variables
func LoadConfig() (*ConfigManager, error) {
	// Get the user's preferred editor, otherwise assign default program
	editor := os.Getenv(envEditor)
	if editor == "" {
		fmt.Println("Defaulting EDITOR to `xdg-open`")
		editor = "xdg-open" // TODO: this might be different for various OS's
	}

	// Get the user's Obsidian vault path from the environment variable
	vaultPath := os.Getenv(envObsidianVaultPath)
	if vaultPath == "" {
		return nil, fmt.Errorf("Environment variable %s is not set", envObsidianVaultPath)
	}

	// Validate the obtained vaultPath
	if err := validateVaultPath(vaultPath); err != nil {
		return nil, err
	}
	dotObsidianPath := filepath.Join(vaultPath, ".obsidian")

	// Import the `daily-notes` plugin json settings from respective json file
	dailyNotesConfig, err := obsidian.LoadDailyNotesConfig(vaultPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load config for daily-notes plugin: %v", err)
	}

	// Import the new file location configuration settings from JSON
	newFileConfig, err := obsidian.LoadNewFileConfig(vaultPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load config for new files: %v", err)
	}

	// Final ConfigManager variable to return before running any config-related commands
	config := &ConfigManager{
		Editor:			editor,

		ObsidianVaultPath:	vaultPath,
		DotObsidianVaultPath:	dotObsidianPath,

		DailyNotesConfig:	dailyNotesConfig,
		NewFileConfig:		newFileConfig,
	}

	return config, nil
}

// validateVaultPath: Validates the user's given obsidian vault path:
// 1. An existing path in the filesystem which is indeed a directory
// 2. A valid Obsidian vault directory (contains a .obsidian/ dir inside)
func validateVaultPath(path string) error {
	// 1. Check if path exists and is a directory
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Path %s does not exist", path)
		}
		return err // Handle any other error (ie. permissions etc.)
	}
	if !info.IsDir() {
		return fmt.Errorf("Path %s is not a directory", path)
	}

	// 2. Check that the directory is a valid obsidian directory (has .obsidian/)
	dotObsidianPath := filepath.Join(path, ".obsidian")
	info, err = os.Stat(dotObsidianPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Path %s is not a valid obsidian directory (missing .obsidian/)", path)
		}
		return err // Handle any other error (ie. permissions etc.)
	}
	if !info.IsDir() {
		return fmt.Errorf(".obsidian is not a directory in %s", path)
	}

	return nil
}
