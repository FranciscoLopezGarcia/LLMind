package main

import "errors"

func AddProject(config Config, project Project) (Config, error) {
	if project.Name == "" {
		return config, errors.New("project name is empty")
	}

	if project.Path == "" {
		return config, errors.New("project path is empty")
	}

	status := ValidateProject(project)
	if status != "ok" {
		return config, errors.New(status)
	}

	for _, existingProject := range config.Project {
		if existingProject.Name == project.Name {
			return config, errors.New("project already exists: " + project.Name)
		}
	}

	config.Project = append(config.Project, project)

	return config, nil
}
func LinkAgentToProject(config Config, projectName string, agentName string, defaultModel string) (Config, error) {
	projectIndex := -1
	for i, project := range config.Project {
		if project.Name == projectName {
			projectIndex = i
			break
		}
	}

	if projectIndex == -1 {
		return config, errors.New("project not found: " + projectName)
	}

	agentFound := false
	for _, agent := range config.Agent {
		if agent.Name == agentName {
			agentFound = true
			break
		}
	}

	if !agentFound {
		return config, errors.New("agent not found: " + agentName)
	}

	for _, linkedAgent := range config.Project[projectIndex].Agents {
		if linkedAgent.AgentName == agentName {
			return config, errors.New("agent already linked to project: " + agentName)
		}
	}

	config.Project[projectIndex].Agents = append(config.Project[projectIndex].Agents, ProjectAgent{
		AgentName:    agentName,
		DefaultModel: defaultModel,
		Enabled:      true,
	})

	return config, nil
}
