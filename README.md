# MCP LSP Go

A Model Context Protocol (MCP) server that allows AI assistants like Claude to interact with Go's Language Server Protocol (LSP) and benefit from advanced Go code analysis features.

## Overview

This MCP server helps AI assistants to:

- Use LSP to analyze Go code
- Navigate to definitions and find references
- Check code diagnostics
- Get hover information for symbols
- Get completion suggestions

## Architecture

This project uses the [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) library to implement the Model Context Protocol. The MCP integration enables seamless communication between AI assistants and Go tools.

The server communicates with [gopls](https://github.com/golang/tools/tree/master/gopls), the official language server for Go, via the Language Server Protocol (LSP).

## Features

- **LSP Integration**: Connection to Go's Language Server Protocol for code analysis
- **Code Navigation**: Finding definitions and references in the code
- **Code Quality**: Getting diagnostics and errors
- **Advanced Information**: Hover information and completion suggestions

## Project Structure

```bash
.
├── cmd
│   └── mcp-gopls        # Application entry point
├── pkg
│   ├── lsp             # LSP client to communicate with gopls
│   │   ├── client      # LSP client implementation
│   │   └── protocol    # LSP protocol types and features
│   ├── server          # MCP server
│   └── tools           # MCP tools exposing LSP features
```

## Installation

```bash
go install github.com/hloiseaufcms/mcp-gopls/cmd/mcp-gopls@latest
```

## Add to Cursor

```json
{
  "mcpServers": {
    "mcp-gopls": {
      "command": "mcp-gopls"
    }
  }
} 
```

## MCP Tools

The MCP server provides the following tools:

| Tool | Description |
|-------|-------------|
| `go_to_definition` | Navigate to the definition of a symbol |
| `find_references` | Find all references to a symbol |
| `check_diagnostics` | Get diagnostics for a file |
| `get_hover_info` | Get detailed information about a symbol |
| `get_completion` | Get completion suggestions at a position |
| `analyze_coverage` | Analyze test coverage for Go code |

## Usage Example

Using the server with AI assistants that support MCP:

```Markdown
# Ask the AI to get information about the code
Can you find the definition of the `ServeStdio` function in this project?

# Ask for diagnostics
Are there any errors in my main.go file?

# Ask for information about a symbol
What does the Context.WithTimeout function do in Go?
```

## Development

```bash
git clone https://github.com/hloiseaufcms/mcp-gopls.git
cd mcp-gopls
go mod tidy
go build -o mcp-gopls cmd/mcp-gopls/main.go
./mcp-gopls
```

## Prerequisites

- Go 1.21 or higher
- gopls installed (`go install golang.org/x/tools/gopls@latest`)

## Integration with Ollama

This MCP server can be used with any tool that supports the MCP protocol. For Ollama integration:

1. Make sure Ollama is running
2. The MCP server runs independently and communicates through stdin/stdout
3. Configure your client to use the MCP server as a tool provider

## License

Apache License 2.0