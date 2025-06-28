package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Init(basePath string) error {
	files := map[string]any{
		"auth/login.lavapi.json": testFileSample(),
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
	return nil

}

func testFileSample() map[string]any {
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
