package main

import (
	"aidd_mcp/tools"
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "aidd-mcp-server",
		Version: "1.0.0",
	}, nil)

	// --- SDD project tools ---

	mcp.AddTool(server, &mcp.Tool{
		Name: "setup_sdd_project",
		Description: "One-time setup: creates the <project>-agent-docs repo with the full SDD folder structure and template files, " +
			"then writes AGENTS.md and CLAUDE.md into the main project repo pointing to the docs repo path.",
	}, tools.HandleSetupSDDProject)

	mcp.AddTool(server, &mcp.Tool{
		Name: "get_project_skill",
		Description: "Loads a project-local Agent Skill (SKILL.md) from supported agent directories " +
			"such as .claude/skills, .cursor/skills, .agents/skills, or .github/skills. " +
			"Returns the full markdown content of the matching skill.",
	}, tools.HandleGetProjectSkill)

	mcp.AddTool(server, &mcp.Tool{
		Name: "write_doc",
		Description: "Writes or updates a markdown spec artifact in the agent-docs repo " +
			"(spec.md, design.md, tasks.md, evidence.md, or any constitution file), " +
			"then commits and pushes. Acquires a lockfile to prevent concurrent write conflicts.",
	}, tools.HandleWriteDoc)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_doc",
		Description: "Reads any markdown file from the agent-docs repo by relative path. Pulls latest before reading.",
	}, tools.HandleGetDoc)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_docs",
		Description: "Lists all markdown files in the agent-docs repo, optionally filtered by a path prefix (e.g. '.specs/my-feature').",
	}, tools.HandleListDocs)

	// stdio: el IDE levanta este binario como subprocess y habla JSON-RPC por stdin/stdout
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
