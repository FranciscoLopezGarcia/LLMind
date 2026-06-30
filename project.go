package main

import (
	"errors"
)

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

	for _, existingProject := range config.Projects {
		if existingProject.Name == project.Name {
			return config, errors.New("project already exists: " + project.Name)
		}
	}

	config.Projects = append(config.Projects, project)

	return config, nil
}
