/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "Lavapi",
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
