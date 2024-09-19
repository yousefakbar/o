package cli

import (
	"fmt"
	"github.com/yousefakbar/o/internal/config"
)

var SearchCommand = Command{
	Name:		"search",
	Aliases:	[]string{"s"},
	RequiresConfig:	true,
	Run:		runCommandSearch,
}

func runCommandSearch(cfg *config.ConfigManager, args []string) error {
	fmt.Println("Searching and opening note in EDITOR")
	// TODO: Implement the actual logic here
	return nil
}
