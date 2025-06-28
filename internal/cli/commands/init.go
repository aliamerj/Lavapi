package commands

import (
	"path/filepath"
	"time"

	"github.com/aliamerj/lavapi/internal/cli/ui"
	"github.com/aliamerj/lavapi/internal/core"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new API test folder with examples",
	Args:  cobra.MaximumNArgs(1),
	Run:  runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	ui.PrintColorfulLogo()
	s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
	s.Prefix = "Initializing project... "
	s.Color("cyan")
	s.Start()
	basePath := filepath.Join(args[0], "lavapi")
	if err := core.Init(basePath); err != nil {
		ui.LogError("", err)
    return;
	}

	s.Stop()
  ui.LogSuccess("Project initialized successfully at " + basePath)
}
