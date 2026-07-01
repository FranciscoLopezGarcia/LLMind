package main

type Config struct {
	Projects []Project `json:"projects"`
	Agents   []Agent   `json:"agents"`
}

type Project struct {
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	DefaultAgent string   `json:"default_agent"`
	Tags         []string `json:"tags"`
}

type Agent struct {
	Name     string   `json:"name"`
	Provider string   `json:"provider"`
	Command  string   `json:"command"`
	Models   []string `json:"models"`
}
