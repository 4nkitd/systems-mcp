package toolsets

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Memory handles reading from and writing to the memory file.
type Memory struct {
	filePath string
}

// NewMemory creates a new Memory instance.
// If no path is provided, it defaults to ~/.mcp/memory.json.
func NewMemory(path string) (*Memory, error) {
	var err error
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		path = filepath.Join(home, ".mcp", "memory.json")
	}

	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create memory directory: %w", err)
	}

	return &Memory{filePath: path}, nil
}

// loadData reads the memory file and unmarshals it into a map.
func (m *Memory) loadData() (map[string]string, error) {
	data := make(map[string]string)
	file, err := os.ReadFile(m.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, return empty map, not an error.
			return data, nil
		}
		return nil, fmt.Errorf("failed to read memory file: %w", err)
	}

	// if the file is empty, no need to unmarshal
	if len(file) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal memory data: %w", err)
	}

	return data, nil
}

// saveData marshals the map and writes it to the memory file.
func (m *Memory) saveData(data map[string]string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal memory data: %w", err)
	}
	if err := os.WriteFile(m.filePath, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write memory file: %w", err)
	}
	return nil
}

// Set stores a key-value pair in the memory file.
func (m *Memory) Set(key, value string) error {
	data, err := m.loadData()
	if err != nil {
		return err
	}
	data[key] = value
	return m.saveData(data)
}

// Get retrieves a value for a given key from the memory file.
// If the key does not exist, it returns an error.
func (m *Memory) Get(key string) (string, error) {
	data, err := m.loadData()
	if err != nil {
		return "", err
	}

	value, ok := data[key]
	if !ok {
		return "", fmt.Errorf("no information found for key: %s", key)
	}
	return value, nil
}

// GetAll retrieves all key-value pairs from the memory file as a JSON string.
func (m *Memory) GetAll() (string, error) {
	data, err := m.loadData()
	if err != nil {
		return "", err
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal memory data: %w", err)
	}
	return string(bytes), nil
}
