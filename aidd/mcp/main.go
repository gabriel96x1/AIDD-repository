package main

import (
	"aidd_mcp/tools"
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	if err := os.MkdirAll(tools.ConventionsDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create conventions directory: %v", err)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "aidd-mcp-server",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_available_sdd_conventions",
		Description: "Dynamically scans the server files and lists all languages/frameworks that currently have coding convention guidelines configured.",
	}, tools.HandleListSDDConventions)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_sdd_coding_conventions",
		Description: "Retrieves the specific markdown style guide content for a requested language or framework.",
	}, tools.HandleGetSDDConventions)

	// stdio: el IDE levanta este binario como subprocess y habla JSON-RPC por stdin/stdout
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
