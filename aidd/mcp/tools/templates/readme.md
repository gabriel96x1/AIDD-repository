# {{PROJECT_NAME}} — Agent Docs

SDD specification repository for **{{PROJECT_NAME}}**.

All specs, designs, tasks, and evidence artifacts live here.
This repo is accessed by coding agents via the AIDD MCP server.

## Structure

```
.specs/
  _constitution/       # Non-negotiable project decisions
    mission.md
    tech-stack.md
    roadmap.md
    coding-conventions/
      index.md
  <feature-slug>/      # One folder per feature
    spec.md
    design.md
    tasks.md
    evidence.md
```

## Navigation

Agents access this repo via MCP tools: `get_doc`, `write_doc`, `list_docs`.
Humans can browse and edit directly — every file is plain Markdown.

## Adding a feature

1. Create `.specs/<feature-slug>/` folder.
2. Use the templates in `.specs/_feature-template/` as starting point.
3. Follow the four-phase SDD loop: spec → design → tasks → evidence.