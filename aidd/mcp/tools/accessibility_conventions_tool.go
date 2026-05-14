package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Accessibility Conventions Directory
const AccessibilityConventionsDir = "../docs/accesibility"

type ListAccessibilityConventionsInput struct{}

func HandleListAccessibilityConventions(ctx context.Context, req *mcp.CallToolRequest, input ListAccessibilityConventionsInput) (*mcp.CallToolResult, string, error) {
	files, err := os.ReadDir(AccessibilityConventionsDir)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read accessibility conventions directory: %w", err)
	}

	var standards []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			standardName := strings.TrimSuffix(file.Name(), ".md")
			standards = append(standards, standardName)
		}
	}

	if len(standards) == 0 {
		return nil, "No accessibility convention files found in the directory.", nil
	}

	response := "Available accessibility standards:\n- " + strings.Join(standards, "\n- ")
	return nil, response, nil
}

type AccessibilityConventionInput struct {
	Standard string `json:"standard" jsonschema:"The accessibility standard name (matches the filename without .md)."`
}

func HandleGetAccessibilityConventions(ctx context.Context, req *mcp.CallToolRequest, input AccessibilityConventionInput) (*mcp.CallToolResult, string, error) {
	cleanStandard := filepath.Base(input.Standard)
	filePath := filepath.Join(AccessibilityConventionsDir, fmt.Sprintf("%s.md", cleanStandard))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Sprintf("Error: Accessibility standard '%s' not found.", input.Standard), nil
	}

	return nil, string(content), nil
}
