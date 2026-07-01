package main

type Config struct {
	Project []Project `json:"project"`
	Agent   []Agent   `json:"agent"`
}

type Project struct {
	Name         string         `json:"name"`
	Path         string         `json:"path"`
	DefaultAgent string         `json:"default_agent"`
	Tags         []string       `json:"tags"`
	Agents       []ProjectAgent `json:"agents"`
}

type Agent struct {
	Name     string   `json:"name"`
	Provider string   `json:"provider"`
	Command  string   `json:"command"`
	Models   []string `json:"models"`
}

type ProjectAgent struct {
	AgentName    string `json:"agent_name"`
	DefaultModel string `json:"default_model"`
	Enabled      bool   `json:"enabled"`
}
