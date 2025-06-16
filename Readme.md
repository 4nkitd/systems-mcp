# Systems MCP Server

A Model Context Protocol (MCP) server that provides system interaction tools for AI assistants. This server enables AI models to interact with your local system through a comprehensive set of tools including volume control, text-to-speech, file system operations, weather information, and more.

## Features

### 🔊 Audio & Volume Control
- **Volume Up/Down**: Adjust system volume by 10 units
- **Mute/Unmute**: Control system audio muting
- **Text-to-Speech**: Convert text to speech using system TTS

### 💾 Memory & Data Management
- **Save Information**: Persist information under a key/value pair (stored in JSON file)
- **Get Saved Info**: Retrieve information by key (stored in JSON file)

### ⏰ Reminders & Alarms
- **Set Alarm**: Schedule alarms with custom messages
- Support for 24-hour time format (HH:MM)

### 📁 File System Operations
- **Current Directory**: Get current working directory
- **List Directory**: Browse directory contents
- **Read Files**: Read file contents from the system

### 🌤️ Internet & Location Services
- **Weather Information**: Get weather data for any location
- **Current Location**: Retrieve current geographical location

## Installation

### Releases

Download the latest prebuilt binaries for your platform from the GitHub Releases page:

- Linux / macOS (tar.gz): https://github.com/4nkitd/systems-mcp/releases/latest/download/paytring_mcp.tar.gz
- Windows (zip): https://github.com/4nkitd/systems-mcp/releases/latest/download/paytring_mcp.zip

After downloading, extract and run:

```bash
tar -xzf paytring_mcp.tar.gz
./paytring_mcp serve
```

For Windows:

```powershell
Expand-Archive paytring_mcp.zip
.\paytring_mcp.exe serve
```


### Prerequisites
- Go 1.23.4 or later
- macOS (primary support)

### Build from Source
```bash
git clone https://github.com/4nkitd/systems-mcp.git
cd systems-mcp/mcp-server
go mod download
go build -o 4nkitd-mcp ./main.go
```

## Usage

### Command Line Options

```bash
./4nkitd-mcp serve [flags]
```

#### Available Flags:
- `--transport`: Transport type (`stdio` or `sse`) - Default: `stdio`
- `--log_dir`: Log directory path - Default: current directory
- `--host`: Host to bind server (SSE only) - Default: `localhost`
- `--port`: Port to bind server (SSE only) - Default: `8080`

### Transport Modes

#### STDIO Mode (Default)
Best for direct integration with AI assistants like Claude Desktop:
```bash
./4nkitd-mcp serve --transport stdio
```

#### SSE Mode
For web-based integrations:
```bash
./4nkitd-mcp serve --transport sse --host localhost --port 8080
```

## Configuration

### Claude Desktop Integration

Add to your Claude Desktop configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "4nkitd-mcp": {
      "command": "/path/to/4nkitd-mcp",
      "args": ["serve", "--transport", "stdio"],
      "env": {}
    }
  }
}
```

### Environment Variables

The server can be configured using the following environment variables:
- `LOG_DIR`: Override default log directory
- `TRANSPORT`: Set transport mode
- `HOST`: Set server host (SSE mode)
- `PORT`: Set server port (SSE mode)

## Available Tools

| Tool Name | Description | Parameters |
|-----------|-------------|------------|
| `volumeUp` | Increase system volume by 10 | None |
| `volumeDown` | Decrease system volume by 10 | None |
| `volumeMute` | Mute system volume | None |
| `volumeUnmute` | Unmute system volume | None |
| `speak` | Text-to-speech conversion | `message` (string) |
| `saveInfo` | Save information to memory | `key` (string), `value` (string) |
| `getSavedInfo` | Retrieve saved information | `key` (string, optional) |
| `setAlarm` | Set an alarm reminder | `time` (HH:MM), `message` (optional) |
| `getCurrentWorkingDirectory` | Get current directory | None |
| `listDirectory` | List directory contents | `path` (optional) |
| `readFile` | Read file contents | `path` (required) |
| `getWeather` | Get weather information | `location` (optional) |
| `getCurrentLocation` | Get current location | None |

## Development

### Project Structure
```
mcp-server/
├── main.go                    # Application entry point
├── cmd/                       # CLI commands and configuration
│   ├── cli.go                # Root command setup
│   ├── serve.go              # Server command implementation
│   ├── config.go             # Configuration management
│   └── constants.go          # Application constants
├── internal/
│   ├── mcp/                  # MCP server implementation
│   │   ├── mcp.go           # Server initialization
│   │   ├── server.go        # Hook registration
│   │   └── tools.go         # Tool registration
│   ├── toolsets/            # Tool implementations
│   │   ├── volume_tools.go  # Audio/volume tools
│   │   ├── memory_tools.go  # Memory management
│   │   ├── reminder_tools.go # Alarm/reminder tools
│   │   ├── filesystem_tools.go # File system operations
│   │   └── internet_tools.go # Weather/location tools
│   └── log/                 # Logging utilities
```

### Adding New Tools

1. Implement the tool function in the appropriate toolset file
2. Register the tool in `internal/mcp/tools.go`
3. Follow the MCP tool specification for parameter definitions

Example:
```go
// In toolsets/example_tools.go
func ExampleTool(ctx context.Context, arguments mcp.CallToolRequestArguments) (*mcp.CallToolResult, error) {
    // Tool implementation
}

// In internal/mcp/tools.go
p.Mcp.AddTool(mcp.NewTool("exampleTool",
    mcp.WithDescription("Description of the tool"),
    mcp.WithString("param", mcp.Description("Parameter description")),
),
    toolsets.ExampleTool,
)
```

## Logging

The server provides comprehensive logging with different levels:
- **DEBUG**: Detailed execution information
- **INFO**: General operational messages
- **ERROR**: Error conditions

Logs are written to the specified log directory or current directory if not specified.

## Security Considerations

- File system tools operate within the permissions of the running user
- Network tools may require internet connectivity
- Consider running with restricted permissions in production environments

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions:
- Open an issue on GitHub
- Check the [MCP specification](https://modelcontextprotocol.io/) for protocol details

## Changelog

### v0.1.0
- Initial release
- Basic tool set implementation
- STDIO and SSE transport support
- Comprehensive logging system
