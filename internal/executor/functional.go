package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aliamerj/lavapi/internal/config"
	"github.com/aliamerj/lavapi/internal/utils"
	"github.com/aliamerj/lavapi/internal/validation"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	cGreen   = color.New(color.FgHiGreen)
	cRed     = color.New(color.FgHiRed)
	cCyan    = color.New(color.FgHiCyan)
	cYellow  = color.New(color.FgHiYellow)
	cGray    = color.New(color.FgHiBlack)
	cBlue    = color.New(color.FgHiBlue)

	symbolPass    = cGreen.Sprint("✔")
	symbolFail    = cRed.Sprint("✖")
	symbolFile    = cYellow.Sprint("📄")
)

func showSpinner(label string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + cCyan.Sprintf(label)
	s.Color("cyan")
	s.Start()
	return s
}

func printFileTree(path string, depth int) {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if i == len(parts)-1 {
			// Last part (file)
			fmt.Printf("%s%s %s\n", strings.Repeat("   ", depth), "└──", cYellow.Sprint(part))
		} else {
			// Directory part
			fmt.Printf("%s%s %s/\n", strings.Repeat("   ", depth), "├──", cBlue.Sprint(part))
			depth++
		}
	}
}

func RunFunctionalTests(file utils.Test, mainConfig *config.Config) (int, int, time.Duration, error) {
	endpoint := strings.TrimSpace(file.API.Endpoint)
	if endpoint == "" {
		return 0, 0, 0, fmt.Errorf("missing endpoint in %s", file.Path)
	}
	if mainConfig != nil {
		endpoint = strings.TrimRight(mainConfig.BaseURL, "/") + endpoint
	}

	// Print file tree structure
	relPath := strings.TrimPrefix(file.Path, "lavapi/")
	printFileTree(relPath, 0)

	var passed, failed int
	var totalDuration time.Duration

	for testName, t := range file.API.Tests.Functional {
		method := strings.ToUpper(t.Method)

		// Start spinner
		label := fmt.Sprintf("%s [%s %s]", testName, method, endpoint)
		spin := showSpinner(label)

		bodyBytes, _ := json.Marshal(t.Body)
		req, err := http.NewRequest(method, endpoint, bytes.NewReader(bodyBytes))
		if err != nil {
			spin.Stop()
			fmt.Printf("   %s %s  %s\n", symbolFail, cRed.Sprintf(testName), cRed.Sprintf("Request build failed"))
			failed++
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: 10 * time.Second}
		start := time.Now()
		resp, err := client.Do(req)
		elapsed := time.Since(start)
		totalDuration += elapsed
		spin.Stop()

		if err != nil {
			fmt.Printf("   %s %s  %s\n", symbolFail, cRed.Sprintf(testName), cRed.Sprintf(err.Error()))
			failed++
			continue
		}

		respBody, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		expectedStatus, hasStatus := t.Expect["status"].(float64)
		if hasStatus && int(expectedStatus) != resp.StatusCode {
			fmt.Printf("   %s %s  %s (expected %s, got %s)  %s\n",
				symbolFail,
				cRed.Sprintf(testName),
				cRed.Sprintf("Status mismatch"),
				cGreen.Sprintf("%v", expectedStatus),
				cRed.Sprintf("%v", resp.StatusCode),
				cGray.Sprintf("(%v)", elapsed.Truncate(time.Millisecond)),
			)
			failed++
			continue
		}

		if err := validation.ValidateExpectations(respBody, t.Expect); err != nil {
			fmt.Printf("   %s %s  %s  %s\n",
				symbolFail,
				cRed.Sprintf(testName),
				cRed.Sprintf(err.Error()),
				cGray.Sprintf("(%v)", elapsed.Truncate(time.Millisecond)),
			)
			failed++
			continue
		}

		fmt.Printf("   %s %s  %s\n",
			symbolPass,
			cGreen.Sprintf(testName),
			cGray.Sprintf("(%v)", elapsed.Truncate(time.Millisecond)))
		passed++
	}

	return passed, failed, totalDuration.Truncate(time.Microsecond), nil
}
