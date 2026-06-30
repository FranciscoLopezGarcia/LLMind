package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome 2 LLMind")
	fmt.Println("===================")

	if len(os.Args) > 1 && os.Args[1] == "init" {
		err := CreateLLMindDirs()
		if err != nil {
			fmt.Println("Error creating LLMind directories:", err)
			return
		}

		err = CreateDefaultConfigFile()
		if err != nil {
			fmt.Println("Error creating default config:", err)
			return
		}

		llmindDir, err := GetLLMindDir()
		if err != nil {
			fmt.Println("Error getting LLMind directory:", err)
			return
		}

		configPath, err := GetDefaultConfigPath()
		if err != nil {
			fmt.Println("Error getting config path:", err)
			return
		}

		fmt.Println()
		fmt.Println("LLMind initialized.")
		fmt.Println("Directory:", llmindDir)
		fmt.Println("Config:", configPath)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "project" {
		if len(os.Args) > 2 && os.Args[2] == "add" {
			if len(os.Args) < 5 {
				fmt.Println("Usage: llmind project add <name> <path>")
				return
			}

			projectName := os.Args[3]
			projectPath := os.Args[4]

			configPath, err := GetDefaultConfigPath()
			if err != nil {
				fmt.Println("Error getting config path:", err)
				return
			}

			config, err := LoadConfig(configPath)
			if err != nil {
				fmt.Println("Error loading config:", err)
				return
			}

			project := Project{
				Name:         projectName,
				Path:         projectPath,
				DefaultAgent: "",
				Tags:         []string{},
			}

			config, err = AddProject(config, project)
			if err != nil {
				fmt.Println("Error adding project:", err)
				return
			}

			err = SaveConfig(configPath, config)
			if err != nil {
				fmt.Println("Error saving config:", err)
				return
			}

			fmt.Println("Project added:", project.Name)
			return
		}
	}
	if len(os.Args) > 1 && os.Args[1] == "agent" {
		if len(os.Args) > 2 && os.Args[2] == "add" {
			if len(os.Args) < 7 {
				fmt.Println("Usage: llmind agent add <name> <provider> <command> <models>")
				fmt.Println("Example: llmind agent add my-agent provider-name agent-cmd model-a,model-b")
				return
			}

			agentName := os.Args[3]
			agentProvider := os.Args[4]
			agentCommand := os.Args[5]
			agentModels := strings.Split(os.Args[6], ",")

			configPath, err := GetDefaultConfigPath()
			if err != nil {
				fmt.Println("Error getting config path:", err)
				return
			}

			config, err := LoadConfig(configPath)
			if err != nil {
				fmt.Println("Error loading config:", err)
				return
			}

			agent := Agent{
				Name:     agentName,
				Provider: agentProvider,
				Command:  agentCommand,
				Models:   agentModels,
			}

			config, err = AddAgent(config, agent)
			if err != nil {
				fmt.Println("Error adding agent:", err)
				return
			}

			err = SaveConfig(configPath, config)
			if err != nil {
				fmt.Println("Error saving config:", err)
				return
			}

			fmt.Println("Agent added:", agent.Name)
			return
		}
	}

	configPath, err := GetDefaultConfigPath()
	if err != nil {
		fmt.Println("Error getting config path:", err)
		return
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Println()
	fmt.Println("Projects:")

	for _, project := range config.Projects {
		fmt.Println("- " + project.Name)
		fmt.Println("  path:", project.Path)
		fmt.Println("  default agent:", project.DefaultAgent)
		fmt.Println("  tags:", strings.Join(project.Tags, ", "))
		fmt.Printf("  status: %s\n", ValidateProject(project))
	}

	fmt.Println()
	fmt.Println("Agents:")

	for _, agent := range config.Agents {
		fmt.Println("- " + agent.Name)
		fmt.Println("  provider:", agent.Provider)
		fmt.Println("  command:", agent.Command)
		fmt.Println("  models:", strings.Join(agent.Models, ", "))
		fmt.Printf("  status: %s\n", ValidateAgents(agent))
	}
}
