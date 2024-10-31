package cli

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/yousefakbar/o/internal/config"
)

var BrowseCommand = Command{
	Name:		"browse",
	Aliases:	[]string{"b"},
	RequiresConfig:	true,
	Run:		runCommandBrowse,
}

func runCommandBrowse(cfg *config.ConfigManager, args []string) error {
	err := launchFileBrowser(cfg.Terminal, cfg.FileBrowser, cfg.ObsidianVaultPath)
	if err != nil {
		return fmt.Errorf("Failed to launch file browser: %w", err)
	}

	return nil
}

func launchFileBrowser(terminal, browser, vaultPath string) error {
	if terminal == "" {
		return fmt.Errorf("Cannot run %s because $TERMINAL is not set", browser)
	}

	cmd := exec.Command(terminal, "-e", browser, vaultPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed running file browser: %w", err)
	}

	return nil
}
