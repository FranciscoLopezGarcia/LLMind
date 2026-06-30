# LLMind

`LLMind` is a local developer tool for managing projects and LLM agents from the terminal.

It provides a lightweight control layer to organize AI/LLM command-line tools and connect them to specific development projects.

Configuration is stored locally in YAML, and LLMind validates project paths and agent commands before using them.

---

## What it does

`LLMind` helps you:

- Manage local projects by name and path.
- Register LLM agents with provider, command, and models.
- Validate project paths and commands available in `PATH`.
- Store configuration in a human-readable YAML file.
- List projects and agents from the terminal.

---

## Key features

- Initialize a local configuration directory in `~/.llmind`
- Save settings in `~/.llmind/config.yaml`
- Add local projects with name and path
- Add LLM agents with provider, command, and models
- Validate existing project paths
- Validate agent commands in the system `PATH`
- Keep configuration editable in YAML
- List configured projects and agents

---

## Requirements

- Go 1.25 or newer

---

## Installation

```bash
git clone <repo-url>
cd LLMind
```

Run locally:

```bash
go run .
```

Build binary:

```bash
go build -o llmind
```

Run binary:

```bash
./llmind
```

---

## Basic usage

### Initialize LLMind

```bash
go run . init
```

This creates the local configuration structure:

```text
~/.llmind/
├── config.yaml
├── logs/
└── tasks/
```

Initial config example:

```yaml
projects: []
agents: []
```

### Add a project

```bash
go run . project add <name> <path>
```

Example:

```bash
go run . project add LLMind /path/to/LLMind
```

This entry is stored in `~/.llmind/config.yaml` as:

```yaml
projects:
  - name: LLMind
    path: /path/to/LLMind
    default_agent: ""
    tags: []
agents: []
```

### Add an agent

```bash
go run . agent add <name> <provider> <command> <models>
```

Example:

```bash
go run . agent add my-agent provider-name agent-cmd model-a,model-b
```

This registers an agent and checks that the command exists in the current `PATH`.

Example YAML:

```yaml
agents:
  - name: my-agent
    provider: provider-name
    command: agent-cmd
    models:
      - model-a
      - model-b
```

### Show configuration

```bash
go run .
```

Example output:

```text
Welcome to LLMind
================

Projects:
- LLMind
  path: /path/to/LLMind
  default agent:
  tags:
  status: ok

Agents:
- my-agent
  provider: provider-name
  command: agent-cmd
  models: model-a, model-b
  status: ok
```

---

## Configuration file

The main configuration file is stored at:

```text
~/.llmind/config.yaml
```

Full example:

```yaml
projects:
  - name: LLMind
    path: /path/to/LLMind
    default_agent: ""
    tags: []

agents:
  - name: my-agent
    provider: provider-name
    command: agent-cmd
    models:
      - model-a
      - model-b
```

> This file is human-readable and can be edited manually, but the preferred workflow is to update it using LLMind commands.

---

## Project structure

```text
LLMind/
├── agent.go
├── config.go
├── go.mod
├── go.sum
├── main.go
├── models.go
├── project.go
├── validate.go
├── examples/
│   └── config.yaml
└── README.md
```

### File responsibilities

- `main.go`: entry point and command routing.
- `config.go`: loads, saves, and initializes configuration.
- `project.go`: project-related logic.
- `agent.go`: agent-related logic.
- `models.go`: core data models.
- `validate.go`: path and command validation.

---

## Notes

- This repository is designed as a local control tool for developers.
- Keep configuration readable and easy to edit.
- Prefer using LLMind commands rather than editing YAML manually.

---

## Safety notes

LLMind is intended to run locally and manage commands that may interact with the developer machine.

Principles:

- Keep configuration local by default.
- Do not store API keys in logs.
- Do not print sensitive environment variables.
- Validate paths before using them.
- Validate commands before running them.
- Avoid destructive actions without explicit confirmation.

---

## Data model

### Config

```go
type Config struct {
    Projects []Project `yaml:"projects"`
    Agents   []Agent   `yaml:"agents"`
}
```

### Project

```go
type Project struct {
    Name         string   `yaml:"name"`
    Path         string   `yaml:"path"`
    DefaultAgent string   `yaml:"default_agent"`
    Tags         []string `yaml:"tags"`
}
```

### Agent

```go
type Agent struct {
    Name     string   `yaml:"name"`
    Provider string   `yaml:"provider"`
    Command  string   `yaml:"command"`
    Models   []string `yaml:"models"`
}
```
