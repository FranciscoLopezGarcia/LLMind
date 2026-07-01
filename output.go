package main

import (
	"fmt"
	"strings"
)

func PrintWelcome() {
	fmt.Println("Welcome 2 LLMind")
	fmt.Println("===================")
}

func PrintHelp() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  llmind init")
	fmt.Println("  llmind list")
	fmt.Println("  llmind tui")
	fmt.Println("  llmind project add <name> <path>")
	fmt.Println("  llmind agent add <name> <provider> <command> <models>")
	fmt.Println("  llmind help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  llmind project add LLMind /home/flg/Escritorio/stuff/LLMind")
	fmt.Println("  llmind agent add claude anthropic claude sonnet,opus")
}

func PrintConfig(config Config) {
	fmt.Println()
	fmt.Println("Project:")

	if len(config.Project) == 0 {
		fmt.Println("  No projects configured.")
	} else {
		for _, project := range config.Project {
			fmt.Println("- " + project.Name)
			fmt.Println("  path:", project.Path)
			fmt.Println("  default agent:", project.DefaultAgent)
			fmt.Println("  tags:", strings.Join(project.Tags, ", "))
			fmt.Printf("  status: %s\n", ValidateProject(project))
			fmt.Println("  linked agents:")
			if len(project.Agents) == 0 {
				fmt.Println("    none")
			} else {
				for _, linkedAgent := range project.Agents {
					fmt.Printf(
						"    - %s model=%s enabled=%t\n",
						linkedAgent.AgentName,
						linkedAgent.DefaultModel,
						linkedAgent.Enabled,
					)
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("Agent:")

	if len(config.Agent) == 0 {
		fmt.Println("  No agents configured.")
	} else {
		for _, agent := range config.Agent {
			fmt.Println("- " + agent.Name)
			fmt.Println("  provider:", agent.Provider)
			fmt.Println("  command:", agent.Command)
			fmt.Println("  models:", strings.Join(agent.Models, ", "))
			fmt.Printf("  status: %s\n", ValidateAgent(agent))
		}
	}
}
