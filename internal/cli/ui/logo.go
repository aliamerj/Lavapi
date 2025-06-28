package ui

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func PrintColorfulLogo() {
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
