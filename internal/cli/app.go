package cli

import (
	"fmt"
	"github.com/yousefakbar/o/internal/config"
	"github.com/yousefakbar/o/internal/utils"
)

type Command struct {
	Name		string
	Aliases		[]string
	RequiresConfig	bool
	Run		func(*config.ConfigManager, []string) error
}

var CommandRegistry = []Command{
	HelpCommand,
	SearchCommand,
	TodayCommand,
	VersionCommand,
	NewCommand,
}

// Run: Executes the main running block of `o` with context of `cfgManager`
func Run(args []string) error {
	// If no command provided, print helpful message and exit with error
	if len(args) < 2 {
		fmt.Println("Usage: o <command> [options]")
		return fmt.Errorf("No command provided (Hint: try 'o help')")
	}

	// Find the requested command based on the user arg
	command	:= args[1]
	cmd	:= FindCommand(command)
	if cmd == nil {
		return fmt.Errorf("Unknown command: %s", command)
	}

	// Create a ConfigManager and load configuration settings if needed
	var cfgManager *config.ConfigManager
	var err error
	if cmd.RequiresConfig {
		cfgManager, err = config.LoadConfig()
		if err != nil {
			return err
		}
	}

	// Run the selected command and return gracefully
	cmd.Run(cfgManager, args)
	return nil
}

func FindCommand(name string) *Command {
	for _, cmd := range CommandRegistry {
		if cmd.Name == name || utils.Contains(cmd.Aliases, name) {
			return &cmd
		}
	}
	return nil
}
