package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/yousefakbar/o/internal/config"
	"github.com/yousefakbar/o/internal/utils"
)

var NewCommand = Command{
	Name:		"new",
	Aliases:	[]string{"n"},
	RequiresConfig:	true,
	Run:		runCommandNew,
}

func runCommandNew(cfg *config.ConfigManager, args []string) error {
	var noteFile string

	// If user passed note name as arg, use that name. Otherwise prompt user for note name
	if len(args) > 2 {
		noteFile = fmt.Sprintf("%s.md", strings.TrimSpace(args[2]))
	} else {
		noteName, err := utils.Prompt("Enter note name (without .md): ")
		if err != nil {
			return fmt.Errorf("Failed to get input for new note name: %w", err)
		}
		noteFile = fmt.Sprintf("%s.md", strings.TrimSpace(noteName))
	}

	// Validate note name
	if noteFile == ".md" {
		return fmt.Errorf("Error: note name cannot be empty")
	}

	notePath := filepath.Join(cfg.NewFileConfig.FolderPath, noteFile)

	// Validate that note file is not already existing
	if _, err := os.Stat(notePath); err == nil {
		return fmt.Errorf("Note file already exists at: %s", notePath)
	}

	// Launch preferred editor for user to edit the new note
	err := launchEditor(cfg.Editor, notePath)
	if err != nil {
		return fmt.Errorf("Failed to launch editor to edit new file: %w", err)
	}

	return nil
}
