package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"github.com/yousefakbar/o/internal/config"
)

var TodayCommand = Command{
	Name:		"today",
	Aliases:	[]string{"t"},
	RequiresConfig:	true,
	Run:		runCommandToday,
}

func runCommandToday(cfg *config.ConfigManager, args []string) error {
	// Convert the user's set daily notes file format to a Go date format (ie. YYYY = 2006 etc.)
	today := time.Now().Format(convertObsidianDateFormat(cfg.DailyNotesConfig.Format))
	todayNotePath := filepath.Join(cfg.ObsidianVaultPath, cfg.DailyNotesConfig.Folder, fmt.Sprintf("%s.md", today))

	err := launchEditor(cfg.Editor, todayNotePath)
	if err != nil {
		return fmt.Errorf("Failed to open today's note: %w", err)
	}

	return nil
}

func launchEditor(editor string, filePath string) error {
	cmd := exec.Command(editor, filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to launch editor: %w", err)
	}

	return nil
}

// TODO: Completely cover MomentJS date formats
// Ref: https://momentjs.com/docs/#/displaying/format/
func convertObsidianDateFormat(format string) string {
	format = strings.Replace(format, "YYYY", "2006", 1)
	format = strings.Replace(format, "MM", "01", 1)
	format = strings.Replace(format, "DD", "02", 1)
	return format
}
