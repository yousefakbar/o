package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"github.com/yousefakbar/o/internal/config"
)

var SearchCommand = Command{
	Name:		"search",
	Aliases:	[]string{"s"},
	RequiresConfig:	true,
	Run:		runCommandSearch,
}

// runCommandSearch Runs the "Search" command for this program
func runCommandSearch(cfg *config.ConfigManager, args []string) error {
	// Capture all .md files using fd
	files, err := findNotesInVault(cfg.ObsidianVaultPath)
	if err != nil {
		return fmt.Errorf("Failed to launch find command: %w", err)
	}

	// Pass list of files to FZF to prompt user to select from the list
	selectFile, err := selectFileWithFzf("fzf", files, cfg.ObsidianVaultPath)
	if err != nil {
		return fmt.Errorf("Failed to launch fzf command: %w", err)
	}

	// Open selected file in user's preferred editor of choice
	err = launchEditor(cfg.Editor, cfg.ObsidianVaultPath + "/" + selectFile)
	if err != nil {
		return fmt.Errorf("Failed to launch editor to edit selected file: %w", err)
	}

	return nil
}

// findNotesInVault Uses find command to capture all .md files in the vault
func findNotesInVault(findDir string) ([]string, error) {
	findCmd := "fd"

	// Verify that the findCmd binary is present in the user PATH
	if _, err := exec.LookPath(findCmd); err != nil {
		findCmd = "find"
	}

	var out bytes.Buffer

	// Run `findCmd` to search and return list of markdown files in findDir
	cmd := exec.Command(findCmd, "-e", "md", ".", findDir)
	if findCmd == "find" {
		// find /home/yha/Documents/YousefBrain/**/*.md
		cmd = exec.Command(findCmd, findDir, "-name", "*.md")
	}

	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Failed to find files: %w", err)
	}

	// Strip the findDir path to make it relative, and convert to list of strings
	fileList := bytes.Split(out.Bytes(), []byte("\n"))
	files, err := stripVaultPath(fileList, findDir)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert filepath to relative path: %w", err)
	}

	return files, nil
}

// selectFileWithFzf Uses FZF to allow the user to select one file from a list of files
func selectFileWithFzf(fzfCmd string, files []string, vaultPath string) (string, error) {
	// Verify that the fzfCmd binary is present in the user PATH
	if _, err := exec.LookPath(fzfCmd); err != nil {
		return "", fmt.Errorf("Failed to find `%s` in PATH", fzfCmd)
	}

	var out bytes.Buffer

	cmd := exec.Command(fzfCmd, "--height", "50%", "--border", "--ansi",
	"--prompt", "🟣 Select a file: ", "--preview", "fzf-preview.sh " +
	vaultPath + "/{}")
	cmd.Stdin = bytes.NewReader([]byte(strings.Join(files, "\n")))
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to run %s with error: %w", fzfCmd, err)
	}

	selectFile := strings.TrimSpace(out.String())

	return selectFile, nil
}

// stripVaultPath Filters the buffer output of a find command, stripping the vaultPath and converting to a slice of strings
func stripVaultPath(fileList [][]byte, vaultPath string) ([]string, error) {
	var files []string
	for _, file := range fileList {
		if len(file) > 0 {
			// Convert absolute file path to relative path (without vault path)
			relPath, err := filepath.Rel(vaultPath, string(file))
			if err != nil {
				return nil, fmt.Errorf("Failed to get relative path: %w", err)
			}
			files = append(files, relPath)
		}
	}

	return files, nil
}
