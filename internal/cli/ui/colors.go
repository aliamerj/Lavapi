package ui

import (
	"github.com/fatih/color"
)

func LogSuccess(message string) {
	color.New(color.Bold, color.FgHiGreen).Printf("✅ %s\n", message)
}

func LogAlert(message string) {
	color.New(color.Bold, color.FgYellow).Printf("⚠️  %s\n", message)
}

func LogError(message string, err error) {
	color.New(color.Bold, color.FgRed).Printf("❌ %s: %v\n", message, err)
}
func LogBadRequest(message string) {
	color.New(color.Bold, color.FgHiRed).Printf("❌ %s \n", message)
}
