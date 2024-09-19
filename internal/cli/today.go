package cli

import (
	"fmt"
	"github.com/yousefakbar/o/internal/config"
	"os"
	"os/exec"
	"time"
)

var TodayCommand = Command{
	Name:		"today",
	Aliases:	[]string{"t"},
	RequiresConfig:	true,
	Run:		runCommandToday,
}

func runCommandToday(cfg *config.ConfigManager, args []string) error {
	todayNotePath := genTodayNotePath(cfg.DailyNotesPath)

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

func genTodayNotePath(dailyNotesPath string) string {
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s/%s.md", dailyNotesPath, today)
}
