package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== ECS CLI Monitor ===")
	fmt.Println("Version: 0.1.0")
	fmt.Println()

	// Check for flags/commands
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "list-services":
		listServices()
	case "list-tasks":
		listTasks()
	case "monitor":
		monitorMode()
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		showHelp()
	}
}

// showHelp displays available commands
func showHelp() {
	help := `
ECS Monitor - CLI tool to monitor AWS ECS services

USAGE:
    ecs-monitor [COMMAND]

COMMANDS:
    list-services     List all ECS services
    list-tasks        List running tasks
    monitor          Interactive monitoring mode (like k9s)
    help             Show this help message

EXAMPLES:
    ecs-monitor list-services
    ecs-monitor monitor
    ecs-monitor help
`
	fmt.Println(help)
}

// listServices displays ECS services
func listServices() {
	fmt.Println("Fetching ECS services...")
	// TODO: Implement AWS SDK integration
	fmt.Println("Services:")
	fmt.Println("  - service-1")
	fmt.Println("  - service-2")
}

// listTasks displays running tasks
func listTasks() {
	fmt.Println("Fetching ECS tasks...")
	// TODO: Implement AWS SDK integration
	fmt.Println("Tasks:")
	fmt.Println("  - task-1 (RUNNING)")
	fmt.Println("  - task-2 (PENDING)")
}

// monitorMode starts interactive monitoring
func monitorMode() {
	fmt.Println("Starting monitor mode...")
	fmt.Println("Press 'q' to quit")
	// TODO: Implement interactive UI with tcell/tview
}
