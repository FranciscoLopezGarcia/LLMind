package main

import (
	"fmt"
	"strings"
)

func HandleInit() {
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
}

func HandleList() {
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

	PrintConfig(config)
}

func HandleProjectCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: llmind project add <name> <path>")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: llmind project add <name> <path>")
			return
		}

		projectName := args[1]
		projectPath := args[2]

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

	default:
		fmt.Println("Unknown project command:", args[0])
		fmt.Println("Usage: llmind project add <name> <path>")
	}
	switch args[0] {
	case "add":
		// lo que ya tenés

	case "link-agent":
		if len(args) < 4 {
			fmt.Println("Usage: llmind project link-agent <project-name> <agent-name> <default-model>")
			return
		}

		projectName := args[1]
		agentName := args[2]
		defaultModel := args[3]

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

		config, err = LinkAgentToProject(config, projectName, agentName, defaultModel)
		if err != nil {
			fmt.Println("Error linking agent:", err)
			return
		}

		err = SaveConfig(configPath, config)
		if err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Println("Agent linked to project:", agentName, "->", projectName)

	default:
		// unknown command
	}
}

func HandleAgentCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: llmind agent add <name> <provider> <command> <models>")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 5 {
			fmt.Println("Usage: llmind agent add <name> <provider> <command> <models>")
			fmt.Println("Example: llmind agent add claude anthropic claude sonnet,opus")
			return
		}

		agentName := args[1]
		agentProvider := args[2]
		agentCommand := args[3]
		agentModels := strings.Split(args[4], ",")

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

	default:
		fmt.Println("Unknown agent command:", args[0])
		fmt.Println("Usage: llmind agent add <name> <provider> <command> <models>")
	}

}
func HandleTUI() {
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

	err = RunTUI(config)
	if err != nil {
		fmt.Println("Error running TUI:", err)
		return
	}
}
