package commands

import (
	"os"
	"path/filepath"

	"github.com/aliamerj/lavapi/internal/cli/ui"
	"github.com/aliamerj/lavapi/internal/config"
	"github.com/aliamerj/lavapi/internal/core"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <path> [flags]",
	Short: "Run Lavapi tests from a folder or selected files",
	Long: `Run tests from a folder of .lavapi.json files.

Examples:
  lavapi run lavapi/ -a
  lavapi run lavapi/ -p auth/login.lavapi.json users/
`,
	Args: cobra.MaximumNArgs(1),
	Run:  runRun,
}

var (
	flagAll      bool
	flagPaths    []string
	flagFailFast bool
)

func init() {
	runCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Run all .lavapi.json tests")
	runCmd.Flags().StringSliceVarP(&flagPaths, "paths", "p", nil, "Run specific test files (relative to test folder)")
	runCmd.Flags().BoolVarP(&flagFailFast, "fail-fast", "f", false, "Stop execution on first test failure")
}

func runRun(cmd *cobra.Command, args []string) {
	if !flagAll && len(flagPaths) == 0 {
		ui.LogBadRequest("You must provide either --all (-a) or --paths (-p)")
		cmd.Help()
		return
	}
	if flagAll && len(flagPaths) > 0 {
		ui.LogBadRequest("You cannot use both --all (-a) and --paths (-p) at the same time")
		cmd.Help()
		return
	}

	lavapiFolder := "lavapi"
	if len(args) == 1 {
		lavapiFolder = args[0]
	}
	info, err := os.Stat(lavapiFolder)
	if err != nil || !info.IsDir() {
		ui.LogError("Test folder not found: "+lavapiFolder, err)
		return
	}

	color.New(color.Bold, color.FgCyan).Printf("🔍 Running Lavapi tests in: %s\n", lavapiFolder)

	configPath := filepath.Join(lavapiFolder, "config.json")

	mainConfig, err := config.ReadConfig(configPath)
	if err != nil {
		ui.LogError("Failed to read "+lavapiFolder+":", err)
		return
	}

	if mainConfig == nil {
		ui.LogAlert("config.json not found, continuing without it...")
	} else {
		ui.LogSuccess("Load config.json")
	}

	if err := core.Run(lavapiFolder, mainConfig, flagAll, flagPaths, flagFailFast); err != nil {
		ui.LogError("", err)
		return
	}
}
