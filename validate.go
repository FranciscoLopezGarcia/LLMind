package main

import (
	"os"
	"os/exec"
)

func ValidateProject(project Project) string {
	path := project.Path

	if path == "" {
		return "Project path is empty"
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "Project path does not exist: " + path
	}
	if err != nil {
		return "Error accessing project path: " + err.Error()
	}
	if !info.IsDir() {
		return "Project path is not a directory: " + path
	}

	return "ok"
}

func ValidateAgent(agent Agent) string {
	if agent.Command == "" {
		return "Agent command is empty"
	}

	_, err := exec.LookPath(agent.Command)
	if err != nil {
		return "Agent command not found: " + agent.Command
	}

	return "ok"
}
