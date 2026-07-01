package main

import (
	"fmt"
	"os"
)

func main() {
	PrintWelcome()

	if len(os.Args) < 2 {
		PrintHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		HandleInit()

	case "list":
		HandleList()

	case "project":
		HandleProjectCommand(os.Args[2:])

	case "agent":
		HandleAgentCommand(os.Args[2:])

	case "help":
		PrintHelp()

	default:
		fmt.Println("Unknown command:", command)
		PrintHelp()
	}
}
