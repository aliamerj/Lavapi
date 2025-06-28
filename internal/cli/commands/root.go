package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lavapi",
	Short: "Lavapi - CLI for API testing and monitoring",
	Long: `Lavapi is a configuration-driven CLI tool for testing API functionality,
performance, and SLA monitoring — all in one.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
}
