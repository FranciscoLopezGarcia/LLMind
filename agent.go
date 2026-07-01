package main

import "errors"

func AddAgent(config Config, agent Agent) (Config, error) {
	if agent.Name == "" {
		return config, errors.New("agent name is empty")
	}

	if agent.Provider == "" {
		return config, errors.New("agent provider is empty")
	}

	if agent.Command == "" {
		return config, errors.New("agent command is empty")
	}

	if len(agent.Models) == 0 {
		return config, errors.New("agent models are empty")
	}

	status := ValidateAgent(agent)
	if status != "ok" {
		return config, errors.New(status)
	}

	for _, existingAgent := range config.Agent {
		if existingAgent.Name == agent.Name {
			return config, errors.New("agent already exists: " + agent.Name)
		}
	}

	config.Agent = append(config.Agent, agent)

	return config, nil
}
