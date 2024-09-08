package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	O_VERSION          = "v0.1"
	obsidianVaultEnv   = "OBSIDIAN_VAULT_PATH"
	journalDir         = "journal"
	defaultEditorEnv   = "EDITOR"
	defaultVaultPath   = "~/Documents/YousefBrain"
)

func assertObsidianVaultPathSet() {
	if os.Getenv(obsidianVaultEnv) == "" {
		fmt.Println("ERROR: OBSIDIAN_VAULT_PATH is not set to vault location.")
		os.Exit(1)
	}
}

func openTodayJournalFile() {
	assertObsidianVaultPathSet()

	todayDate := time.Now().Format("2006-01-02")
	vaultPath := os.Getenv(obsidianVaultEnv)
	if vaultPath == "" {
		vaultPath = defaultVaultPath
	}
	todayNoteFilename := filepath.Join(vaultPath, journalDir, todayDate+".md")

	editor := os.Getenv(defaultEditorEnv)
	if editor == "" {
		fmt.Println("ERROR: $EDITOR environment variable is not set.")
		os.Exit(1)
	}

	cmd := exec.Command(editor, todayNoteFilename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("ERROR: Failed to open file %s: %v\n", todayNoteFilename, err)
		os.Exit(1)
	}
}

func searchMdFilesWithFzf() {
	assertObsidianVaultPathSet()

	if _, err := exec.LookPath("fzf"); err != nil {
		fmt.Println("ERROR: fzf is not installed.")
		os.Exit(2)
	}

	vaultPath := os.Getenv(obsidianVaultEnv)
	if vaultPath == "" {
		vaultPath = defaultVaultPath
	}

	cmd := exec.Command("fd", "--glob", "*.md", vaultPath)
	fdOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("ERROR: Failed to find .md files: %v\n", err)
		os.Exit(2)
	}

	fzfCmd := exec.Command("fzf", "--preview", "bat --style=grid --color=always --line-range=:500 {}", "--preview-window=right:60%", "--height=40%", "--border", "--ansi", "--delimiter", "/", "--nth=-1", "--prompt=🟣 ")
	fzfCmd.Stdin = strings.NewReader(string(fdOutput))
	fzfOutput, err := fzfCmd.Output()
	if err != nil {
		fmt.Printf("ERROR: Failed to select file with fzf: %v\n", err)
		os.Exit(2)
	}
	selectedFilePath := strings.TrimSuffix(string(fzfOutput), "\n")

	editor := os.Getenv(defaultEditorEnv)
	if editor == "" {
		fmt.Println("ERROR: $EDITOR environment variable is not set.")
		os.Exit(1)
	}

	editCmd := exec.Command(editor, selectedFilePath)
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr
	editCmd.Stdin = os.Stdin
	if err := editCmd.Run(); err != nil {
		fmt.Printf("ERROR: Failed to open file %s: %v\n", string(fzfOutput), err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [subcommand]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h", "--help", "help", "h":
		fmt.Println("Help information for the script")
	case "-s", "--search", "search", "s":
		searchMdFilesWithFzf()
	case "-t", "--today", "today", "t":
		openTodayJournalFile()
	case "-v", "--version", "version", "v":
		fmt.Printf("o: %s\n", O_VERSION)
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		fmt.Println("Usage: go run main.go [subcommand]")
		os.Exit(1)
	}
}
