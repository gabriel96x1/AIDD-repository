# Coding Conventions Index — {{PROJECT_NAME}}

The agent reads this file first for any implementation task.

Identify which skills apply to the current task and load them via the
MCP tool `get_project_skill` before writing any code.

This system is agent-agnostic and supports skills generated or installed
by tools such as autoskills, gh skill, Claude Code, Codex, Cursor,
Copilot, and compatible Agent Skills implementations.

---

## Project Skills

| Scope | Skill name | When to load |
|---|---|---|
| {{LANGUAGE_OR_FRAMEWORK_1}} | `{{skill-1}}` | {{e.g. Any time you write or modify source files in this language}} |
| {{LANGUAGE_OR_FRAMEWORK_2}} | `{{skill-2}}` | {{e.g. Any time you create or modify UI components}} |
| {{CONCERN_1}} | `{{skill-3}}` | {{e.g. Any time you write or modify tests}} |

> Use the MCP tool:
>
> `get_project_skill(main_repo_path="{{MAIN_REPO_PATH}}", skill_name="{{skill-1}}")`
>
> to retrieve the full SKILL.md content for a given skill.

---

## Skill Resolution

The MCP server automatically resolves skills from supported agent locations
inside the main project repository.

Search priority:

1. `.claude/skills/<skill>/SKILL.md`
2. `.cursor/skills/<skill>/SKILL.md`
3. `.agents/skills/<skill>/SKILL.md`
4. `.github/skills/<skill>/SKILL.md`

This allows the same SDD workflow to work across multiple coding agents
without coupling the project to a single vendor or directory structure.

---

## Non-Negotiable Rules (All Files)

These apply regardless of language, framework, or loaded skill.

1. Every file has one responsibility.
2. No commented-out code committed to main.
3. No `TODO` comments without a linked issue or task reference.
4. All magic numbers extracted to named constants.
5. All public functions have explicit types on inputs and outputs.
6. No implicit `any` (or language equivalent) without a documented reason.
7. Prefer composition over inheritance unless explicitly justified.
8. Keep business logic separated from UI and infrastructure concerns.
9. Avoid hidden side effects and mutable shared state.
10. New code must align with the existing project architecture and loaded skills.

---

## Installing New Skills

Skills should be installed into the main project repository using any
compatible Agent Skills tooling.

Examples include:

- autoskills
- gh skill
- Claude Code skills
- Cursor skills
- Copilot Agent Skills

Once installed, the MCP server automatically discovers them through
`get_project_skill`.

No changes to the MCP server are required when adding project-local skills.

---

## Adding a Skill to This Index

1. Install or create the skill in the main project repository.
2. Add an entry to the table above.
3. Commit with message:

   `chore(skills): add {{skill-name}} skill`

4. The agent will automatically load it in future sessions when relevant.

---

## Example

| Scope | Skill name | When to load |
|---|---|---|
| React Native | `react-native` | Any time you create screens, hooks, or RN components |
| TypeScript | `typescript` | Any time you modify TS source files |
| Testing | `testing` | Any time you write unit or integration tests |

Example tool call:

```text
get_project_skill(
  main_repo_path="/Users/user/projects/my-app",
  skill_name="react-native"
)