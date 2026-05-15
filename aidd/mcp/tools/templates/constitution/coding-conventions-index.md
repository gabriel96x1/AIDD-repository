# Coding Conventions Index — {{PROJECT_NAME}}

The agent reads this file first for any implementation task.
Identify which conventions apply and fetch the corresponding document
via `get_sdd_coding_conventions` before writing any code.

---

## Convention Documents

| Scope | Convention key | When to read |
|---|---|---|
| {{LANGUAGE_OR_FRAMEWORK_1}} | `{{key-1}}` | {{e.g. Any time you write or modify source files in this language}} |
| {{LANGUAGE_OR_FRAMEWORK_2}} | `{{key-2}}` | {{e.g. Any time you create or modify UI components}} |
| {{CONCERN_1}} | `{{key-3}}` | {{e.g. Any time you write or modify tests}} |

> Use the MCP tool `get_sdd_coding_conventions(language="{{key-1}}")` to retrieve
> the full convention document for a given key.

---

## Non-Negotiable Rules (All Files)

These apply regardless of language or framework and are not repeated in
individual convention documents.

1. Every file has one responsibility.
2. No commented-out code committed to main.
3. No `TODO` comments without a linked issue or task reference.
4. All magic numbers extracted to named constants.
5. All public functions have explicit types on inputs and outputs.
6. No implicit `any` (or language equivalent) without a documented reason in a comment.

---

## Adding a New Convention Document

1. Add the `.md` file to the conventions directory in the MCP server.
2. Add an entry to the table above via `write_doc`.
3. Commit with message: `chore(conventions): add {{language}} conventions`.
4. The agent will pick it up automatically in the next session via `list_available_sdd_conventions`.