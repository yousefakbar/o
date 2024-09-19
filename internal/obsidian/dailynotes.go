package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
)

// loadDailyNotesConfig: Loads the setting indicating where daily notes are stored (if not default location)
func LoadDailyNotesConfig(vaultPath string) (string, error) {
	// Define struct to hold the daily notes folder configuration value
	var dailyNotesConfig struct {
		Folder string `json:"folder"`
	}

	// Define the json file for the obsidian `daily-notes` plugin settings
	dailyNotesConfigPath := filepath.Join(vaultPath, ".obsidian", "daily-notes.json")

	// Read the JSON setting form dailyNotesConfigPath and store in the struct
	err := ReadJSONSettings(dailyNotesConfigPath, &dailyNotesConfig)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Assuming default daily notes configuration (daily-notes.json does not exist")
			return vaultPath, nil
		}
	}

	// If dailyNotesConfig.Folder setting is missing or empty, return the default which is vaultPath
	if err := ValidateSettingKey(dailyNotesConfig.Folder); err != nil {
		return vaultPath, nil
	}

	return filepath.Join(vaultPath, dailyNotesConfig.Folder), nil
}
