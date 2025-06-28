package validation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
)

func ValidateExpectations(respBody []byte, expect map[string]any) error {
	var actualBody map[string]any
	if err := json.Unmarshal(respBody, &actualBody); err != nil {
		return fmt.Errorf("invalid JSON response: %w", err)
	}

	if bodyExpect, ok := expect["body"]; ok {
		if err := deepMatch(bodyExpect, actualBody, "body"); err != nil {
			return err
		}
	}

	// 2. Handle dot-notation: "body.res", "body.user.email", etc.
	for key, rawExpected := range expect {
		if key == "status" || key == "body" {
			continue
		}

		if !strings.HasPrefix(key, "body.") {
			continue // ignore unsupported keys
		}

		path := strings.TrimPrefix(key, "body.")
		actual := gjson.GetBytes(respBody, path)

		if !actual.Exists() {
			return fmt.Errorf("expected %s to exist, but it doesn't", key)
		}

		switch v := rawExpected.(type) {
		case string:
			if v == "!!exists" {
				continue // already checked exists
			}
			if v == "!!not_exists" {
				return fmt.Errorf("expected %s to NOT exist, but it does", key)
			}
			if actual.String() != v {
				return fmt.Errorf("expected %s to be '%s', got '%s'", key, v, actual.String())
			}
		default:
			if !reflect.DeepEqual(actual.Value(), v) {
				return fmt.Errorf("expected %s to be %v, got %v", key, v, actual.Value())
			}
		}
	}

	return nil
}

func deepMatch(expectPart any, actualPart any, path string) error {
	expectMap, ok1 := expectPart.(map[string]any)
	actualMap, ok2 := actualPart.(map[string]any)

	if !ok1 || !ok2 {
		if !reflect.DeepEqual(expectPart, actualPart) {
			return fmt.Errorf("expected %s to be %v, got %v", path, expectPart, actualPart)
		}
		return nil
	}

	for key, expectedVal := range expectMap {
		actualVal, ok := actualMap[key]
		fullPath := fmt.Sprintf("%s.%s", path, key)

		if !ok {
			return fmt.Errorf("expected key %s to exist, but it doesn't", fullPath)
		}

		if err := deepMatch(expectedVal, actualVal, fullPath); err != nil {
			return err
		}
	}
	return nil
}
