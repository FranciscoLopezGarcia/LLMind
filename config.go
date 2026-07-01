package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// LoadConfig loads LLMind data from a JSON file.
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// GetLLMindDir returns the local LLMind directory: ~/.llmind
func GetLLMindDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".llmind"), nil
}

// GetDefaultConfigPath returns the default JSON config path: ~/.llmind/data.json
func GetDefaultConfigPath() (string, error) {
	llmindDir, err := GetLLMindDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(llmindDir, "data.json"), nil
}

// CreateLLMindDirs creates the local LLMind directory structure.
func CreateLLMindDirs() error {
	llmindDir, err := GetLLMindDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(llmindDir, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(llmindDir, "tasks"), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(llmindDir, "logs"), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(llmindDir, "projects"), 0755); err != nil {
		return err
	}

	return nil
}

// CreateDefaultConfigFile creates ~/.llmind/data.json if it does not exist.
func CreateDefaultConfigFile() error {
	configPath, err := GetDefaultConfigPath()
	if err != nil {
		return err
	}

	_, err = os.Stat(configPath)
	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	defaultConfig := Config{
		Projects: []Project{},
		Agents:   []Agent{},
	}

	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// SaveConfig saves LLMind data to a JSON file.
func SaveConfig(path string, config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
