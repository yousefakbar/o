package obsidian

import (
	"fmt"
	"os"
	"path/filepath"
)

// Define struct to hold the daily notes folder configuration value
type DailyNotesConfig struct {
	Folder string `json:"folder"`
	Format string `json:"format"`
}

// loadDailyNotesConfig: Loads the setting indicating where daily notes are stored (if not default location)
func LoadDailyNotesConfig(vaultPath string) (*DailyNotesConfig, error) {
	var dailyNotesConfig DailyNotesConfig

	// Define the json file for the obsidian `daily-notes` plugin settings
	dailyNotesConfigPath := filepath.Join(vaultPath, ".obsidian", "daily-notes.json")

	// Read the JSON setting form dailyNotesConfigPath and store in the struct
	err := ReadJSONSettings(dailyNotesConfigPath, &dailyNotesConfig)
	if err != nil {
		if os.IsNotExist(err) {
			// Set default values for the configuration fields
			fmt.Println("Assuming default daily notes configuration (daily-notes.json does not exist")
			dailyNotesConfig.Folder = vaultPath
			dailyNotesConfig.Format = "YYYY-MM-DD"
			return &dailyNotesConfig, nil
		}
		return nil, fmt.Errorf("Failed to load daily notes configuration file: %v", err)
	}

	// If dailyNotesConfig.Folder setting is missing or empty, return the default which is vaultPath
	if err := ValidateSettingKey(dailyNotesConfig.Folder); err != nil {
		dailyNotesConfig.Folder = vaultPath
	}

	// Validate the format setting; if it's invalid or missing, set to default YYYY-MM-DD
	if err := ValidateSettingKey(dailyNotesConfig.Format); err != nil {
		dailyNotesConfig.Format = "YYYY-MM-DD"
	}

	return &dailyNotesConfig, nil
}
