package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/yousefakbar/o/internal/config"
)

var BrowseCommand = Command{
	Name:		"browse",
	Aliases:	[]string{"b"},
	RequiresConfig:	true,
	Run:		runCommandBrowse,
}

func runCommandBrowse(cfg *config.ConfigManager, args []string) error {
	// Browse and select file using `gum file`
	path, err := launchGumFileBrowser(cfg.ObsidianVaultPath)
	if err != nil {
		return fmt.Errorf("Failed to launch file browser: %w", err)
	}

	// Open selected file in user's preferred editor of choice
	err = launchEditor(cfg.Editor, path)
	if err != nil {
		return fmt.Errorf("Failed to launch editor to edit selected file: %w", err)
	}

	return nil
}

func launchGumFileBrowser(vaultPath string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command(
		"gum",
		"file",
		vaultPath,
		"--height=10",
	)
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Failed running file browser: %w", err)
	}

	selectFile := strings.TrimSpace(out.String())

	return selectFile, nil
}
