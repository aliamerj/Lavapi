package core

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliamerj/lavapi/internal/cli/ui"
	"github.com/aliamerj/lavapi/internal/config"
	"github.com/aliamerj/lavapi/internal/executor"
	"github.com/aliamerj/lavapi/internal/utils"
	"github.com/fatih/color"
)

var (
	cBold      = color.New(color.Bold)
	cGreen     = color.New(color.FgHiGreen)
	cRed       = color.New(color.FgHiRed)
	cCyan      = color.New(color.FgHiCyan)
	cBlue      = color.New(color.FgHiBlue)
	cMagenta   = color.New(color.FgHiMagenta)

	symbolSummary = cMagenta.Sprint("📊")
	symbolTime    = cCyan.Sprint("⏱")
)

func Run(lavapiFolder string, mainConfig *config.Config, flagAll bool, flagPaths []string, flagFailFast bool) error {
	testsFiles, err := findAllTests(lavapiFolder)
	if err != nil {
		return err
	}
	tests, err := utils.ValidateTestFiles(testsFiles)
	if err != nil {
		return err
	}
	fmt.Printf("%s %s/\n", cBlue.Sprint("📁"), color.New(color.FgHiGreen).Sprint("lavapi"))

  var passed, failed int
  var totalTime time.Duration 
  var testFailed error 
  for _, test := range tests {
	 passed, failed, totalTime ,testFailed = executor.RunFunctionalTests(test, mainConfig)
    if testFailed!= nil {
			if flagFailFast {
				return testFailed
			}
			ui.LogAlert("Continuing after failure")
		}
    passed +=passed
    failed +=failed
    totalTime +=totalTime
	}

  	// Summary with colored boxes
	fmt.Printf("\n%s %s %s %s %s %s  \n",
		symbolSummary,
		cBold.Sprint("Summary:"),
		cGreen.Sprintf("✅ %d passed", passed),
		cRed.Sprintf("❌ %d failed", failed),
		symbolTime,
		cCyan.Sprintf("%v", totalTime),
	)

	return nil
}
func findAllTests(folderPath string) ([]string, error) {
	var testFiles []string
	if err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".lavapi.json") {
			testFiles = append(testFiles, path)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("Failed walking test directory: %w", err)
	}

	if len(testFiles) == 0 {
		return nil, fmt.Errorf("No .lavapi.json test files found")
	}
	return testFiles, nil
}
