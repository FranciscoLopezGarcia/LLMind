package main

type Config struct {
	Projects []Project `yaml:"projects"`
	Agents   []Agent   `yaml:"agents"`
}

type Project struct {
	Name         string   `yaml:"name"`
	Path         string   `yaml:"path"`
	DefaultAgent string   `yaml:"default_agent"`
	Tags         []string `yaml:"tags"`
}

type Agent struct {
	Name     string   `yaml:"name"`
	Provider string   `yaml:"provider"`
	Command  string   `yaml:"command"`
	Models   []string `yaml:"models"`
}
