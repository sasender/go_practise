# ECS CLI Monitor

A CLI tool to monitor AWS ECS services and tasks, inspired by k9s (Kubernetes monitor).

## Project Goal

Build an interactive terminal UI to:
- View ECS services and their status
- Monitor running tasks
- Display resource usage (CPU, memory)
- Stream logs from containers
- Navigate with keyboard shortcuts like k9s

## Architecture

```
ecs-monitor/
├── main.go              # CLI entry point and commands
├── aws/                 # AWS ECS API integration
│   ├── client.go        # ECS client setup
│   ├── services.go      # Service operations
│   └── tasks.go         # Task operations
├── ui/                  # Terminal UI (TUI) components
│   ├── dashboard.go     # Main dashboard view
│   ├── list.go          # List rendering
│   └── colors.go        # Color scheme
├── models/              # Data structures
│   ├── service.go       # Service model
│   └── task.go          # Task model
└── config/              # Configuration
    └── config.go        # AWS credentials & settings
```

## Learning Path

### Phase 1: Basic CLI (Current)
- [x] Command-line argument parsing
- [ ] AWS SDK setup
- [ ] Fetch and display services
- [ ] Fetch and display tasks

### Phase 2: Data Models
- [ ] Create Service struct
- [ ] Create Task struct
- [ ] Parse AWS responses

### Phase 3: Formatting & Display
- [ ] Pretty print output
- [ ] Add colors (ANSI)
- [ ] Format timestamps

### Phase 4: Interactive UI
- [ ] Real-time updates
- [ ] Keyboard navigation
- [ ] Interactive filtering

### Phase 5: Advanced Features
- [ ] Log streaming
- [ ] Resource monitoring (CPU, memory)
- [ ] Service actions (restart, scale)

## Setup Instructions

### Prerequisites
- Go 1.21+
- AWS Account with ECS services
- AWS CLI configured with credentials

### Installation

```bash
cd d:\Go_learning\ecs-monitor
go mod download
go run main.go help
```

## Current Features

- Command-line interface with help
- Command routing (list-services, list-tasks, monitor)
- Placeholder implementations ready for AWS integration

## Next Steps

1. Install AWS SDK: `go get -u github.com/aws/aws-sdk-go-v2`
2. Create AWS client to fetch real data
3. Build models for Service and Task
4. Add pretty formatting with colors
5. Implement interactive UI

## Useful Libraries

- **AWS SDK**: `github.com/aws/aws-sdk-go-v2` - Official AWS Go SDK
- **TUI Framework**: `github.com/rivo/tview` - Advanced text-based UI
- **Terminal Colors**: `github.com/fatih/color` - Simple color output
- **Progress**: `github.com/schollz/progressbar` - Progress indicators

## Learning Resources

- [AWS SDK for Go v2](https://aws.github.io/aws-sdk-go-v2/)
- [k9s source code](https://github.com/derailed/k9s) - Great reference!
- [Go command-line apps](https://golang.org/doc/effective_go)

## Tips for Building This

1. **Start simple** - Get basic list-services working first
2. **Add one feature at a time** - Don't build UI before getting data
3. **Test with real AWS** - Use your actual ECS cluster
4. **Reference k9s** - See how they structure similar functionality
5. **Iterative UI** - Start with text output, add UI later
