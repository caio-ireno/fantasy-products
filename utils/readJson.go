package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

func ReadJson[T any](domain string) ([]T, error) {

	// Interpolar o domain no path do arquivo
	filePath := fmt.Sprintf("docs/db/json/%s.json", domain)

	slog.Info("Reading JSON file", "path", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Failed to read file", "error", err, "path", filePath)
		return nil, err
	}

	var result []T
	err = json.Unmarshal(data, &result)
	if err != nil {
		slog.Error("Failed to unmarshal JSON", "error", err)
		return nil, err
	}
	slog.Info("JSON file loaded successfully", "result", result)

	slog.Info("JSON file loaded successfully", "count", len(result))
	return result, nil
}
