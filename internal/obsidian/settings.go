package obsidian

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadJSONSettings: Reads json file `filePath` and unmarshals data to `setting` interface based on requested key
func ReadJSONSettings(filePath string, setting interface{}) error {
	// Check if the JSON filePath exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("File %s does not exist", filePath)
	}

	// Open the JSON file so we can read the contents
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("There was an issue opening file %s: %v", filePath, err)
	}
	defer file.Close()

	// Decode and unmarshal the JSON file contents to the target interface
	if err := json.NewDecoder(file).Decode(&setting); err != nil {
		return fmt.Errorf("There was an issue parsing JSON file %s: %v", filePath, err)
	}

	return nil
}

// ValidateSettingKey: Validates if `settingValue` is indeed set to a value. Used in conjunction with ReadJSONSettings
func ValidateSettingKey(settingValue string) error {
	if settingValue == "" {
		return fmt.Errorf("Setting value is empty or not set")
	}
	return nil
}
