package tools

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ---------------------------------------------------------------------------
// Embedded templates
// Each variable maps 1-to-1 to a .md file in tools/templates/.
// To update a template: edit the .md file and recompile. No Go changes needed.
// ---------------------------------------------------------------------------

//go:embed templates/readme.md
var tmplReadme string

//go:embed templates/constitution/mission.md
var tmplMission string

//go:embed templates/constitution/tech-stack.md
var tmplTechStack string

//go:embed templates/constitution/roadmap.md
var tmplRoadmap string

//go:embed templates/constitution/coding-conventions-index.md
var tmplConventionsIndex string

//go:embed templates/feature/spec.md
var tmplSpec string

//go:embed templates/feature/design.md
var tmplDesign string

//go:embed templates/feature/tasks.md
var tmplTasks string

//go:embed templates/feature/evidence.md
var tmplEvidence string

// ---------------------------------------------------------------------------
// Template rendering
// Only substitutes {{PROJECT_NAME}} and {{DATE}} — the rest of the
// {{PLACEHOLDERS}} are left intact for the agent to fill when drafting specs.
// ---------------------------------------------------------------------------

func renderTemplate(tmpl, projectName string) string {
	r := strings.NewReplacer(
		"{{PROJECT_NAME}}", projectName,
		"{{DATE}}", time.Now().Format("2006-01-02"),
	)
	return r.Replace(tmpl)
}

// ---------------------------------------------------------------------------
// Tool: get_project_skill
// ---------------------------------------------------------------------------

type GetProjectSkillInput struct {
	MainRepoPath string `json:"main_repo_path" jsonschema:"Absolute path to the main project repo."`
	SkillName    string `json:"skill_name"     jsonschema:"Name of the skill folder (e.g. 'react-native', 'typescript')."`
}

type GetProjectSkillOutput struct {
	Content string `json:"content"  jsonschema:"Full SKILL.md content."`
	FoundAt string `json:"found_at" jsonschema:"Relative path where the skill was found."`
}

func HandleGetProjectSkill(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GetProjectSkillInput,
) (*mcp.CallToolResult, GetProjectSkillOutput, error) {

	if input.MainRepoPath == "" || input.SkillName == "" {
		return nil, GetProjectSkillOutput{}, fmt.Errorf("main_repo_path and skill_name are required")
	}

	// Supported agent skill locations (priority order)
	candidates := []string{
		filepath.Join(".claude", "skills", input.SkillName, "SKILL.md"),
		filepath.Join(".cursor", "skills", input.SkillName, "SKILL.md"),
		filepath.Join(".agents", "skills", input.SkillName, "SKILL.md"),
		filepath.Join(".github", "skills", input.SkillName, "SKILL.md"),
	}

	for _, rel := range candidates {
		abs := filepath.Join(input.MainRepoPath, rel)

		data, err := os.ReadFile(abs)
		if err == nil {
			return nil, GetProjectSkillOutput{
				Content: string(data),
				FoundAt: rel,
			}, nil
		}

		if !os.IsNotExist(err) {
			return nil, GetProjectSkillOutput{}, fmt.Errorf(
				"failed reading skill at %s: %w",
				rel,
				err,
			)
		}
	}

	return nil, GetProjectSkillOutput{}, fmt.Errorf(
		"skill '%s' not found in any supported agent directory",
		input.SkillName,
	)
}

// ---------------------------------------------------------------------------
// Tool: setup_sdd_project
// ---------------------------------------------------------------------------

type SetupSDDInput struct {
	ProjectName  string `json:"project_name"    jsonschema:"The name of the project (e.g. 'my-app'). The docs repo will be created as 'my-app-agent-docs'."`
	DocsBasePath string `json:"docs_base_path"  jsonschema:"Absolute path to the directory where the -agent-docs repo will be created (e.g. '/Users/you/projects')."`
	MainRepoPath string `json:"main_repo_path"  jsonschema:"Absolute path to the main project repo where AGENTS.md and CLAUDE.md will be written."`
}

type SetupSDDOutput struct {
	DocsRepoPath string `json:"docs_repo_path" jsonschema:"Absolute path to the created -agent-docs repo."`
	AgentsMDPath string `json:"agents_md_path" jsonschema:"Absolute path to the AGENTS.md written in the main repo."`
	Summary      string `json:"summary"        jsonschema:"Human-readable summary of what was created."`
}

func HandleSetupSDDProject(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input SetupSDDInput,
) (*mcp.CallToolResult, SetupSDDOutput, error) {

	if input.ProjectName == "" || input.DocsBasePath == "" || input.MainRepoPath == "" {
		return nil, SetupSDDOutput{}, fmt.Errorf("project_name, docs_base_path, and main_repo_path are all required")
	}

	docsRepoPath := filepath.Join(input.DocsBasePath, input.ProjectName+"-agent-docs")

	// 1. Create directory structure
	for _, d := range []string{
		filepath.Join(docsRepoPath, ".specs", "_constitution", "coding-conventions"),
		filepath.Join(docsRepoPath, ".specs", "_feature-template"),
	} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return nil, SetupSDDOutput{}, fmt.Errorf("failed to create directory %s: %w", d, err)
		}
	}

	// 2. Write all files from embedded templates
	type entry struct {
		relPath string
		tmpl    string
	}
	for _, f := range []entry{
		{"README.md", tmplReadme},
		{".specs/_constitution/mission.md", tmplMission},
		{".specs/_constitution/tech-stack.md", tmplTechStack},
		{".specs/_constitution/roadmap.md", tmplRoadmap},
		{".specs/_constitution/coding-conventions-index.md", tmplConventionsIndex},
		{".specs/_feature-template/spec.md", tmplSpec},
		{".specs/_feature-template/design.md", tmplDesign},
		{".specs/_feature-template/tasks.md", tmplTasks},
		{".specs/_feature-template/evidence.md", tmplEvidence},
	} {
		absPath := filepath.Join(docsRepoPath, f.relPath)
		if err := os.WriteFile(absPath, []byte(renderTemplate(f.tmpl, input.ProjectName)), 0644); err != nil {
			return nil, SetupSDDOutput{}, fmt.Errorf("failed to write %s: %w", f.relPath, err)
		}
	}

	// 3. Init git repo and first commit
	for _, gitCmd := range [][]string{
		{"init"},
		{"add", "."},
		{"commit", "-m", "chore: initialize SDD structure"},
	} {
		if err := runGit(docsRepoPath, gitCmd...); err != nil {
			return nil, SetupSDDOutput{}, fmt.Errorf("git %s failed: %w", gitCmd[0], err)
		}
	}

	// 4. Write AGENTS.md and CLAUDE.md into the main repo
	// These are built dynamically (not embedded) because they contain docsRepoPath,
	// which is only known at setup time.
	agentsMDPath := filepath.Join(input.MainRepoPath, "AGENTS.md")
	if err := os.WriteFile(agentsMDPath, []byte(buildAgentsMD(input.ProjectName, docsRepoPath)), 0644); err != nil {
		return nil, SetupSDDOutput{}, fmt.Errorf("failed to write AGENTS.md: %w", err)
	}
	claudeMDPath := filepath.Join(input.MainRepoPath, "CLAUDE.md")
	if err := os.WriteFile(claudeMDPath, []byte(buildClaudeMD(input.ProjectName, docsRepoPath)), 0644); err != nil {
		return nil, SetupSDDOutput{}, fmt.Errorf("failed to write CLAUDE.md: %w", err)
	}

	return nil, SetupSDDOutput{
		DocsRepoPath: docsRepoPath,
		AgentsMDPath: agentsMDPath,
		Summary: fmt.Sprintf(
			"SDD setup complete.\n\nDocs repo:  %s\nAGENTS.md:  %s\nCLAUDE.md:  %s\n\n"+
				"Next: fill _constitution/ files using write_doc, starting with mission.md.",
			docsRepoPath, agentsMDPath, claudeMDPath,
		),
	}, nil
}

// ---------------------------------------------------------------------------
// Tool: write_doc
// ---------------------------------------------------------------------------

type WriteDocInput struct {
	DocsRepoPath  string `json:"docs_repo_path"  jsonschema:"Absolute path to the -agent-docs repo."`
	RelativePath  string `json:"relative_path"   jsonschema:"Path relative to docs repo root (e.g. '.specs/my-feature/spec.md')."`
	Content       string `json:"content"         jsonschema:"Full markdown content to write to the file."`
	CommitMessage string `json:"commit_message"  jsonschema:"Git commit message (e.g. 'spec(auth): add login spec')."`
}

type WriteDocOutput struct {
	FilePath  string `json:"file_path"  jsonschema:"Absolute path to the written file."`
	Committed bool   `json:"committed"  jsonschema:"Whether the file was successfully committed and pushed."`
}

func HandleWriteDoc(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input WriteDocInput,
) (*mcp.CallToolResult, WriteDocOutput, error) {

	if input.DocsRepoPath == "" || input.RelativePath == "" || input.Content == "" || input.CommitMessage == "" {
		return nil, WriteDocOutput{}, fmt.Errorf("all fields are required")
	}

	absPath := filepath.Join(input.DocsRepoPath, input.RelativePath)

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return nil, WriteDocOutput{}, fmt.Errorf("failed to create parent directories: %w", err)
	}

	// Lockfile prevents concurrent write+push conflicts between multiple agent instances
	lockPath := filepath.Join(input.DocsRepoPath, ".specs", ".lock")
	if err := acquireLock(lockPath); err != nil {
		return nil, WriteDocOutput{}, fmt.Errorf("repo locked by another operation, retry in a moment: %w", err)
	}
	defer releaseLock(lockPath)

	_ = runGit(input.DocsRepoPath, "pull", "--rebase") // non-fatal if no remote yet

	if err := os.WriteFile(absPath, []byte(input.Content), 0644); err != nil {
		return nil, WriteDocOutput{}, fmt.Errorf("failed to write file: %w", err)
	}

	if err := runGit(input.DocsRepoPath, "add", input.RelativePath); err != nil {
		return nil, WriteDocOutput{}, fmt.Errorf("git add failed: %w", err)
	}
	if err := runGit(input.DocsRepoPath, "commit", "-m", input.CommitMessage); err != nil {
		return nil, WriteDocOutput{}, fmt.Errorf("git commit failed: %w", err)
	}

	_ = runGit(input.DocsRepoPath, "push") // non-fatal if remote not configured yet

	return nil, WriteDocOutput{FilePath: absPath, Committed: true}, nil
}

// ---------------------------------------------------------------------------
// Tool: get_doc
// ---------------------------------------------------------------------------

type GetDocInput struct {
	DocsRepoPath string `json:"docs_repo_path" jsonschema:"Absolute path to the -agent-docs repo."`
	RelativePath string `json:"relative_path"  jsonschema:"Path relative to docs repo root (e.g. '.specs/_constitution/mission.md')."`
}

type GetDocOutput struct {
	Content string `json:"content" jsonschema:"Markdown content of the requested file."`
}

func HandleGetDoc(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GetDocInput,
) (*mcp.CallToolResult, GetDocOutput, error) {

	if input.DocsRepoPath == "" || input.RelativePath == "" {
		return nil, GetDocOutput{}, fmt.Errorf("docs_repo_path and relative_path are required")
	}

	_ = runGit(input.DocsRepoPath, "pull", "--rebase")

	data, err := os.ReadFile(filepath.Join(input.DocsRepoPath, input.RelativePath))
	if err != nil {
		return nil, GetDocOutput{}, fmt.Errorf("file not found: %s", input.RelativePath)
	}

	return nil, GetDocOutput{Content: string(data)}, nil
}

// ---------------------------------------------------------------------------
// Tool: list_docs
// ---------------------------------------------------------------------------

type ListDocsInput struct {
	DocsRepoPath string `json:"docs_repo_path" jsonschema:"Absolute path to the -agent-docs repo."`
	Prefix       string `json:"prefix"         jsonschema:"Optional path prefix to filter results (e.g. '.specs/my-feature'). Leave empty to list all docs."`
}

type ListDocsOutput struct {
	Files []string `json:"files" jsonschema:"List of relative file paths found under the given prefix."`
}

func HandleListDocs(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input ListDocsInput,
) (*mcp.CallToolResult, ListDocsOutput, error) {

	if input.DocsRepoPath == "" {
		return nil, ListDocsOutput{}, fmt.Errorf("docs_repo_path is required")
	}

	_ = runGit(input.DocsRepoPath, "pull", "--rebase")

	root := input.DocsRepoPath
	if input.Prefix != "" {
		root = filepath.Join(root, input.Prefix)
	}

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && info.Name() != ".specs" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".md") {
			rel, _ := filepath.Rel(input.DocsRepoPath, path)
			files = append(files, rel)
		}
		return nil
	})
	if err != nil {
		return nil, ListDocsOutput{}, fmt.Errorf("failed to walk directory: %w", err)
	}

	return nil, ListDocsOutput{Files: files}, nil
}

// ---------------------------------------------------------------------------
// AGENTS.md and CLAUDE.md — built dynamically, NOT embedded
// Reason: they contain docsRepoPath, which is only known at setup time.
// ---------------------------------------------------------------------------

func buildAgentsMD(projectName, docsRepoPath string) string {
	return fmt.Sprintf(`# Agent Operating Rules — %s

This file is the operational constitution for every coding agent working in
this repository. Read it fully before taking any action.

---

## Docs Repository

All SDD artifacts (specs, designs, tasks, evidence, constitution) live in a
**separate repository** accessed via the AIDD MCP server tools.

**Docs repo path:** %s

| Tool | When to use |
|---|---|
| `+"`get_doc`"+` | Read any spec, constitution, or convention file |
| `+"`write_doc`"+` | Create or update any spec artifact |
| `+"`list_docs`"+` | Discover what specs and docs exist |
| `+"`get_project_skill`"+` | Load project-local Agent Skills (Claude, Cursor, Copilot, etc.) |

---

## Before You Do Anything

Read these files via `+"`get_doc`"+` before producing any artifact or code:

1. `+"`"+`.specs/_constitution/mission.md`+"`"+`
2. `+"`"+`.specs/_constitution/tech-stack.md`+"`"+`
3. `+"`"+`.specs/_constitution/roadmap.md`+"`"+`
4. `+"`"+`.specs/_constitution/coding-conventions-index.md`+"`"+`

Then identify which skills apply to the task and load them via:

`+"`get_project_skill(main_repo_path=\"<repo>\", skill_name=\"<skill>\")`"+`

If any file is missing or contains unfilled `+"`{{PLACEHOLDER}}`"+` values,
STOP and NOTIFY the human before proceeding.

ASK for clarifications on the mission, tech stack, roadmap, or conventions if any ambiguity exists. 
NEVER assume or guess on these — they are the source of truth for your work.

---

## The SDD Contract

- NEVER write implementation code without a reviewed `+"`spec.md`"+` AND `+"`design.md`"+`..
- NEVER skip or merge phases: Specify → Plan → Execute → Verify.
- At every gate (G1–G4), stop and state:
  "Gate G[N] reached — awaiting human review of [artifact path]."
- Do not proceed past a gate without explicit human approval.

## Governance Gates

| Gate | Artifact | What human reviews |
|---|---|---|
| G1 | `+"`spec.md`"+` | Goal clarity, EARS criteria, non-goals completeness |
| G2 | `+"`design.md`"+` | Architecture, new dependencies, tech-stack alignment |
| G3 | `+"`tasks.md`"+` | Task atomicity, dependency order, test coverage plan |
| G4 | `+"`evidence.md`"+` | Every criterion has PASS with a verifiable reference |

## Filesystem Rules (This Repo)

- Write source code only when a task in an approved `+"`tasks.md`"+` authorizes it.
- Never modify `+"`AGENTS.md`"+` or `+"`CLAUDE.md`"+` autonomously.
- All spec artifacts go to the docs repo via `+"`write_doc`"+` — never here.

## Ambiguity Protocol

1. Stop immediately.
2. State the ambiguity with options.
3. Wait for human resolution.
4. Record the resolution via `+"`write_doc`"+` before continuing.

## Memory Protocol

Conversation history is not a source of truth. If a decision was made in
chat but not written via `+"`write_doc`"+`, it does not exist.

## Commit Format

`+"```"+`
<type>(<scope>): <short description>

Refs: .specs/<feature-slug>/tasks.md#task-N
Criterion: <paste criterion text>
`+"```"+`

Types: feat, fix, spec, test, refactor, chore.

## Replanning Triggers

After every feature ships, flag if any of the following occurred:
- A tech-stack decision was violated or extended without a constitution update.
- A new pattern was introduced not covered by the constitution.
- The roadmap phase order was disrupted.
- A dependency was added without passing G2.
`, projectName, docsRepoPath)
}

func buildClaudeMD(projectName, docsRepoPath string) string {
	return fmt.Sprintf(`# Claude Agent Rules — %s

This project uses Spec-Driven Development (SDD).

Read `+"`AGENTS.md`"+` in full before taking any action.

**Docs repo path:** %s
Access via: `+"`get_doc`"+`, `+"`write_doc`"+`, `+"`list_docs`"+`.

## Claude-specific notes

- Use extended thinking before any spec artifact or architectural decision.
- Subagents must also read `+"`AGENTS.md`"+` before acting.
- Gate phrasing: "Gate G[N] reached — [path] is ready for your review."
- One clarifying question at a time, never batched.
`, projectName, docsRepoPath)
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, out)
	}
	return nil
}

func acquireLock(lockPath string) error {
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL, 0600)
		if err == nil {
			f.Close()
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("timeout waiting for lock after 10s")
}

func releaseLock(lockPath string) { os.Remove(lockPath) }
