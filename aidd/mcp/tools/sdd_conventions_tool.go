package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SDD Conventions Directory
const SDDConventionsDir = "../docs/sdd"

type ListSDDConventionsInput struct{}

func HandleListSDDConventions(ctx context.Context, req *mcp.CallToolRequest, input ListSDDConventionsInput) (*mcp.CallToolResult, string, error) {
	files, err := os.ReadDir(SDDConventionsDir)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read SDD conventions directory: %w", err)
	}

	var guidelines []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			guideName := strings.TrimSuffix(file.Name(), ".md")
			guidelines = append(guidelines, guideName)
		}
	}

	if len(guidelines) == 0 {
		return nil, "No SDD convention files found in the directory.", nil
	}

	response := "Available SDD guidelines:\n- " + strings.Join(guidelines, "\n- ")
	return nil, response, nil
}

type SDDConventionInput struct {
	Guideline string `json:"guideline" jsonschema:"The SDD guideline name (matches the filename without .md)."`
}

func HandleGetSDDConventions(ctx context.Context, req *mcp.CallToolRequest, input SDDConventionInput) (*mcp.CallToolResult, string, error) {
	cleanGuideline := filepath.Base(input.Guideline)
	filePath := filepath.Join(SDDConventionsDir, fmt.Sprintf("%s.md", cleanGuideline))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Sprintf("Error: SDD guideline '%s' not found.", input.Guideline), nil
	}

	return nil, string(content), nil
}
