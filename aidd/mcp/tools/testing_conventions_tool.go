package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Testing Conventions Directory
const TestingConventionsDir = "../docs/testing"

type ListTestingConventionsInput struct{}

func HandleListTestingConventions(ctx context.Context, req *mcp.CallToolRequest, input ListTestingConventionsInput) (*mcp.CallToolResult, string, error) {
	files, err := os.ReadDir(TestingConventionsDir)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read testing conventions directory: %w", err)
	}

	var strategies []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			strategyName := strings.TrimSuffix(file.Name(), ".md")
			strategies = append(strategies, strategyName)
		}
	}

	if len(strategies) == 0 {
		return nil, "No testing convention files found in the directory.", nil
	}

	response := "Available testing strategies:\n- " + strings.Join(strategies, "\n- ")
	return nil, response, nil
}

type TestingConventionInput struct {
	Strategy string `json:"strategy" jsonschema:"The testing strategy name (matches the filename without .md)."`
}

func HandleGetTestingConventions(ctx context.Context, req *mcp.CallToolRequest, input TestingConventionInput) (*mcp.CallToolResult, string, error) {
	cleanStrategy := filepath.Base(input.Strategy)
	filePath := filepath.Join(TestingConventionsDir, fmt.Sprintf("%s.md", cleanStrategy))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Sprintf("Error: Testing strategy '%s' not found.", input.Strategy), nil
	}

	return nil, string(content), nil
}
