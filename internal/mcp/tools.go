package mcp

import (
	"github.com/4nkitd/systems-mcp/internal/toolsets"
	"github.com/mark3labs/mcp-go/mcp"
)

func (p *ankitd) RegisterTools() {

	// Volume control tools
	p.Mcp.AddTool(mcp.NewTool("volumeUp",
		mcp.WithDescription("Increase system volume by 10"),
	),
		toolsets.VolumeUp,
	)

	p.Mcp.AddTool(mcp.NewTool("volumeDown",
		mcp.WithDescription("Decrease system volume by 10"),
	),
		toolsets.VolumeDown,
	)

	p.Mcp.AddTool(mcp.NewTool("volumeMute",
		mcp.WithDescription("Mute system volume"),
	),
		toolsets.VolumeMute,
	)

	p.Mcp.AddTool(mcp.NewTool("volumeUnmute",
		mcp.WithDescription("Unmute system volume"),
	),
		toolsets.VolumeUnmute,
	)

	p.Mcp.AddTool(mcp.NewTool("speak",
		mcp.WithDescription("Speak text using text-to-speech"),
		mcp.WithString("message", mcp.Description("Text to speak")),
	),
		toolsets.Speak,
	)

	// Memory tools
	p.Mcp.AddTool(mcp.NewTool("saveInfo",
		mcp.WithDescription("Save information to remember"),
		mcp.WithString("key", mcp.Description("Key to store information under")),
		mcp.WithString("value", mcp.Description("Information to store")),
	),
		p.memoryTools.SaveInfo,
	)

	p.Mcp.AddTool(mcp.NewTool("getSavedInfo",
		mcp.WithDescription("Get saved information. If no key is provided, all information is returned."),
		mcp.WithString("key", mcp.Description("Key to retrieve information (optional)")),
	),
		p.memoryTools.GetSavedInfo,
	)

	// Reminder tools
	p.Mcp.AddTool(mcp.NewTool("setAlarm",
		mcp.WithDescription("Set an alarm to remind for tasks"),
		mcp.WithString("time", mcp.Description("Time in HH:MM format (24-hour)")),
		mcp.WithString("message", mcp.Description("Alarm message (optional)")),
	),
		toolsets.SetAlarm,
	)

	// Filesystem tools
	p.Mcp.AddTool(mcp.NewTool("getCurrentWorkingDirectory",
		mcp.WithDescription("Get current working directory"),
	),
		toolsets.GetCurrentWorkingDirectory,
	)

	p.Mcp.AddTool(mcp.NewTool("listDirectory",
		mcp.WithDescription("List directory contents"),
		mcp.WithString("path", mcp.Description("Directory path to list (optional, defaults to current directory)")),
	),
		toolsets.ListDirectory,
	)

	p.Mcp.AddTool(mcp.NewTool("readFile",
		mcp.WithDescription("Read file contents"),
		mcp.WithString("path", mcp.Description("File path to read")),
	),
		toolsets.ReadFile,
	)

	// Internet tools
	p.Mcp.AddTool(mcp.NewTool("getWeather",
		mcp.WithDescription("Get weather information for a location"),
		mcp.WithString("location", mcp.Description("Location to get weather for (optional, defaults to current location)")),
	),
		p.internetTools.GetWeather,
	)

	p.Mcp.AddTool(mcp.NewTool("getCurrentLocation",
		mcp.WithDescription("Get current location information"),
	),
		p.internetTools.GetCurrentLocation,
	)

	p.Mcp.AddTool(mcp.NewTool("fetchURL",
		mcp.WithDescription("Fetch the markdown content of a URL"),
		mcp.WithString("url", mcp.Description("URL to fetch")),
	),
		p.internetTools.FetchURL,
	)
}
