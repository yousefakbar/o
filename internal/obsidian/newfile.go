package obsidian

import (
	"fmt"
	"path/filepath"
)

// "newFileLocation": "current",   e.g. current, root, folder
// "newFileFolderPath": "inbox",   only set with "folder" location
type NewFileConfig struct {
	Location	string `json:"newFileLocation"`
	FolderPath	string `json:"newFileFolderPath"`
}

func LoadNewFileConfig(vaultPath string) (*NewFileConfig, error) {
	var newFileConfig NewFileConfig

	// Define path to the JSON file for new file settings
	newFileConfigPath := filepath.Join(vaultPath, ".obsidian", "app.json")

	// Read the JSON settings and store in the struct
	err := ReadJSONSettings(newFileConfigPath, &newFileConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to load new file configuration on file: %v", err)
	}

	// If Location is not set, default to root vault path
	if err := ValidateSettingKey(newFileConfig.Location); err != nil {
		newFileConfig.Location = "root"
		newFileConfig.FolderPath = vaultPath
		return &newFileConfig, nil
	}

	// If Location is set to "current", default to root vault path
	if newFileConfig.Location == "folder" {
		newFileConfig.FolderPath = filepath.Join(vaultPath, newFileConfig.FolderPath)
	} else {
		newFileConfig.FolderPath = vaultPath
	}

	return &newFileConfig, nil
}
