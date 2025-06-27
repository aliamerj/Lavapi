/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new API test folder with examples",
	Args:  cobra.ExactArgs(1),
	RunE:  Run,
}

func Run(cmd *cobra.Command, args []string) error {
	printColorfulLogo()
	s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
	s.Prefix = "Initializing project... "
	s.Color("cyan")
	s.Start()
	basePath := filepath.Join(args[0], "lavapi")
	files := map[string]any{
		"auth/login.lavapi.json": loginSample(),
		"auth/register.lavapi.json": map[string]string{
			"$schema": "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		},
		"users/create-user.lavapi.json": map[string]string{
			"$schema": "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		},
		"users/get-user.lavapi.json": map[string]string{
			"$schema": "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		},
		"users/update-user.lavapi.json": map[string]string{
			"$schema": "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		},
		"orders/orders.lavapi.json": map[string]string{
			"$schema": "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		},
		"config.json": map[string]string{
			"base_url": "http://localhost:8080",
		},
	}
	for realPath, content := range files {
		fullPath := filepath.Join(basePath, realPath)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create dir: %w", err)
		}
		data, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to write json: %w", err)
		}
		if err := os.WriteFile(fullPath, data, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	s.Stop()
	color.New(color.FgHiGreen, color.Bold).Printf("✅ Project initialized successfully at %s\n", basePath)
	return nil
}

func loginSample() map[string]any {
	return map[string]any{
		"$schema":  "https://aliamerj.github.io/lavapi-json/lavapi.schema.json",
		"endpoint": "/api/auth/login",
		"tests": map[string]any{
			"functional": map[string]any{
				"valid_login": map[string]any{
					"method": "POST",
					"body": map[string]string{
						"email":    "test@example.com",
						"password": "password123",
					},
					"expect": map[string]any{
						"status":          200,
						"body.token":      "exists",
						"body.user.email": "test@example.com",
					},
				},
				"invalid_credentials": map[string]any{
					"method": "POST",
					"body": map[string]string{
						"email":    "wrong@example.com",
						"password": "wrong",
					},
					"expect": map[string]any{
						"status": 401,
					},
				},
			},
		},
	}
}

func printColorfulLogo() {
	colors := []func(a ...any) string{
		color.New(color.FgCyan).SprintFunc(),
		color.New(color.FgHiYellow).SprintFunc(),
		color.New(color.FgHiYellow).SprintFunc(),
		color.New(color.FgRed).SprintFunc(),
		color.New(color.FgRed).SprintFunc(),
	}

	logoLines := []string{
		"██╗      █████╗ ██╗   ██╗ █████╗ ██████╗ ██╗",
		"██║     ██╔══██╗██║   ██║██╔══██╗██╔══██╗██║",
		"██║     ███████║██║   ██║███████║██████╔╝██║",
		"██║     ██╔══██║╚██╗ ██╔╝██╔══██║██╔═══╝ ██║",
		"███████╗██║  ██║ ╚████╔╝ ██║  ██║██║     ██║",
		"╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚═╝  ╚═╝╚═╝     ╚═╝",
	}

	for i, line := range logoLines {
		fmt.Println(colors[i%len(colors)](line))
		time.Sleep(60 * time.Millisecond)
	}

	fmt.Println()
	color.New(color.Bold).Println("⚡ Lavapi 🔥: A Modern CLI for API Testing ✨")
	fmt.Println()
}
