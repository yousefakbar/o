package cli

import (
	"fmt"
	"github.com/yousefakbar/o/internal/config"
)

const Version = "dev"

var VersionCommand = Command{
	Name:		"version",
	Aliases:	[]string{"v"},
	RequiresConfig:	false,
	Run:		runCommandVersion,
}

func runCommandVersion(cfg *config.ConfigManager, args []string) error {
	fmt.Println("o", Version)
	return nil
}
