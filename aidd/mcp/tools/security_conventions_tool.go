package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Security Conventions Directory
const SecurityConventionsDir = "../docs/security"

type ListSecurityConventionsInput struct{}

func HandleListSecurityConventions(ctx context.Context, req *mcp.CallToolRequest, input ListSecurityConventionsInput) (*mcp.CallToolResult, string, error) {
	files, err := os.ReadDir(SecurityConventionsDir)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read security conventions directory: %w", err)
	}

	var practices []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			practiceName := strings.TrimSuffix(file.Name(), ".md")
			practices = append(practices, practiceName)
		}
	}

	if len(practices) == 0 {
		return nil, "No security convention files found in the directory.", nil
	}

	response := "Available security practices:\n- " + strings.Join(practices, "\n- ")
	return nil, response, nil
}

type SecurityConventionInput struct {
	Practice string `json:"practice" jsonschema:"The security practice name (matches the filename without .md)."`
}

func HandleGetSecurityConventions(ctx context.Context, req *mcp.CallToolRequest, input SecurityConventionInput) (*mcp.CallToolResult, string, error) {
	cleanPractice := filepath.Base(input.Practice)
	filePath := filepath.Join(SecurityConventionsDir, fmt.Sprintf("%s.md", cleanPractice))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Sprintf("Error: Security practice '%s' not found.", input.Practice), nil
	}

	return nil, string(content), nil
}
