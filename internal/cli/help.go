package cli

import (
	"fmt"
	"github.com/yousefakbar/o/internal/config"
)

var HelpCommand = Command{
	Name:		"help",
	Aliases:	[]string{"h"},
	RequiresConfig:	false,
	Run:		runCommandHelp,
}

func runCommandHelp(cfg *config.ConfigManager, args []string) error {
	fmt.Println("o -- The Obsidian CLI Wrapper (https://github.com/yousefakbar/o)")
	fmt.Println("Usage: o <command> [options]")
	fmt.Println("")
	fmt.Println("Available Commands (Shorthand, Command):")
	fmt.Printf("  h, help\t\t-- Prints this help message\n")
	fmt.Printf("  n, new\t\t-- Creates a new note\n")
	fmt.Printf("  s, search\t\t-- Searches and opens note in EDITOR\n")
	fmt.Printf("  t, today\t\t-- Opens journal file in EDITOR\n")
	fmt.Printf("  v, version\t\t-- Prints the program version\n")

	return nil
}
