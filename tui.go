package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RunTUI(config Config) error {
	app := tview.NewApplication()

	projectList := tview.NewList()
	projectList.SetBorder(true)
	projectList.SetTitle(" Project ")

	workspaceView := tview.NewTextView()
	workspaceView.SetBorder(true)
	workspaceView.SetTitle(" Workspace ")
	workspaceView.SetDynamicColors(true)
	workspaceView.SetWrap(true)
	workspaceView.SetScrollable(true)

	logView := tview.NewTextView()
	logView.SetBorder(true)
	logView.SetTitle(" Logs ")
	logView.SetDynamicColors(true)
	logView.SetWrap(true)
	logView.SetScrollable(true)

	focusables := []tview.Primitive{
		projectList,
		workspaceView,
		logView,
	}

	focusIndex := 0
	if len(config.Project) == 0 {
		projectList.AddItem("No projects configured", "", 0, nil)
		workspaceView.SetText("No project selected.")
		logView.SetText("Run: llmind project add <name> <path>")
	} else {
		for _, project := range config.Project {
			currentProject := project

			projectList.AddItem(
				currentProject.Name,
				currentProject.Path,
				0,
				func() {
					renderProjectWorkspace(workspaceView, logView, config, currentProject)
				},
			)
		}

		renderProjectWorkspace(workspaceView, logView, config, config.Project[0])
	}

	projectList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(config.Project) {
			renderProjectWorkspace(workspaceView, logView, config, config.Project[index])
		}
	})

	projectList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(config.Project) {
			renderProjectWorkspace(workspaceView, logView, config, config.Project[index])
		}
	})

	layout := tview.NewFlex().
		AddItem(projectList, 0, 1, true).
		AddItem(workspaceView, 0, 2, false).
		AddItem(logView, 0, 2, false)

	app.SetRoot(layout, true)
	app.SetFocus(projectList)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			focusIndex = (focusIndex + 1) % len(focusables)
			app.SetFocus(focusables[focusIndex])
			return nil

		case tcell.KeyEsc:
			app.Stop()
			return nil
		}

		if event.Rune() == 'q' {
			app.Stop()
			return nil
		}

		return event
	})

	return app.Run()
}

func renderProjectWorkspace(workspaceView *tview.TextView, logView *tview.TextView, config Config, project Project) {
	workspaceView.Clear()

	fmt.Fprintf(workspaceView, "[yellow]Project:[white] %s\n", project.Name)
	fmt.Fprintf(workspaceView, "[yellow]Path:[white] %s\n", project.Path)
	fmt.Fprintf(workspaceView, "[yellow]Default agent:[white] %s\n", project.DefaultAgent)
	fmt.Fprintf(workspaceView, "[yellow]Tags:[white] %s\n", strings.Join(project.Tags, ", "))
	fmt.Fprintf(workspaceView, "[yellow]Status:[white] %s\n", ValidateProject(project))

	fmt.Fprintln(workspaceView)
	fmt.Fprintln(workspaceView, "[green]Agents:[white]")

	if len(project.Agents) == 0 {
		fmt.Fprintln(workspaceView, "No agents linked to this project.")
	} else {
		for _, linkedAgent := range project.Agents {
			fmt.Fprintf(
				workspaceView,
				"- %s model=%s enabled=%t\n",
				linkedAgent.AgentName,
				linkedAgent.DefaultModel,
				linkedAgent.Enabled,
			)
		}
	}

	logView.Clear()
	fmt.Fprintf(logView, "[yellow]Project log:[white] %s\n\n", project.Name)
	fmt.Fprintln(logView, "No logs yet.")
	fmt.Fprintln(logView)
	fmt.Fprintln(logView, "Future:")
	fmt.Fprintln(logView, "- orchestrator.log")
	fmt.Fprintln(logView, "- agent-specific logs")
	fmt.Fprintln(logView, "- task/run output")
	fmt.Fprintln(logView)
	fmt.Fprintln(logView, "q: quit")
}
