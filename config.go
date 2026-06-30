package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// carga de configs
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// obtener el directorio ~/.llmind
func GetLLMindDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".llmind"), nil
}

// obetener default path ~/.llmind/config.yaml
func GetDefaultConfigPath() (string, error) {
	llmindDir, err := GetLLMindDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(llmindDir, "config.yaml"), nil
}

// crear ~/.llmind, ~/.llmind/tasks, ~/.llmind/logs
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
	return nil
}

// crear ~/.llmind/config.yaml si noe xiste
func CreateDefaultConfigFile() error {
	configPath, err := GetDefaultConfigPath()
	if err != nil {
		return err
	}
	_, err = os.Stat(configPath)
	if err == nil {
		return nil // El archivo ya existe
	}
	if !os.IsNotExist(err) {
		return err // Otro error al verificar el archivo
	}

	defaultConfig := []byte("projects: []\nagents: []\n")
	return os.WriteFile(configPath, defaultConfig, 0644)
}

// guardar config del input
func SaveConfig(path string, config Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
