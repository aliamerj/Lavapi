package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aliamerj/lavapi/internal/config"
)

type Test struct {
	Path string
	API config.TestFile
}

func ValidateTestFiles(testsFiles []string) ([]Test, error) {
	compiledSchema, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load local schema: %w", err)
	}

	var tests []Test
	for _, file := range testsFiles {
		data, err := os.ReadFile(file)
		if err != nil {
      return nil, fmt.Errorf("Failed to read: %s :\n %w ", file, err)
		}
		var rawTest any
		if err := json.Unmarshal(data, &rawTest); err != nil {
      return nil,	fmt.Errorf("Invalid %s :\n %w ", file, err)

		}

		if err := compiledSchema.Validate(rawTest); err != nil {
      return nil,	fmt.Errorf("Schema validation failed in %s :\n %w ",file, err)
		}

		var test config.TestFile
		if err := json.Unmarshal(data, &test); err != nil {
     return nil, fmt.Errorf("Failed to parse test file %s : %w", file, err)
		}

		tests = append(tests, Test{Path: file, API: test})
	}

  return tests, nil
}
