package mcp

import (
	"github.com/4nkitd/mcp/internal/toolsets"
	"github.com/mark3labs/mcp-go/mcp"
)

func (p *Paytring) RegisterTools() {

	p.Mcp.AddTool(mcp.NewTool("volumeUp",
		mcp.WithDescription("Increse system volume by 10"),
	),
		toolsets.VolumeUp,
	)

	p.Mcp.AddTool(mcp.NewTool("VolumeDown",
		mcp.WithDescription("Descrise system volume by 10"),
	),
		toolsets.VolumeDown,
	)

	p.Mcp.AddTool(mcp.NewTool("VolumeMute",
		mcp.WithDescription("Mute system volume"),
	),
		toolsets.VolumeMute,
	)

	p.Mcp.AddTool(mcp.NewTool("VolumeUnmute",
		mcp.WithDescription("UnMute system volume"),
	),
		toolsets.VolumeUnmute,
	)

	// SaveToInfo
	p.Mcp.AddTool(mcp.NewTool("SaveInfo",
		mcp.WithDescription("SaveInfo to rememeber"),
	),
		toolsets.SaveInfo,
	)

	p.Mcp.AddTool(mcp.NewTool("GetSavedInfo",
		mcp.WithDescription("GetSavedInfo you saved rememeber"),
	),
		toolsets.GetSavedInfo,
	)

	// alarm
	p.Mcp.AddTool(mcp.NewTool("SetAlarm",
		mcp.WithDescription("SetAlarm to remind ankit for stuff"),
	),
		toolsets.SetAlarm,
	)
}
