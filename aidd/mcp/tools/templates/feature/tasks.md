# Tasks: {{Feature Name}}

**Slug:** `{{feature-slug}}`
**Spec:** `.specs/{{feature-slug}}/spec.md`
**Design:** `.specs/{{feature-slug}}/design.md`
**Status:** DRAFT
**Created:** {{DATE}}

---

## Execution Rules

- Work through tasks in dependency order.
- Complete one task fully before starting the next:
  write code → write tests → verify criterion → commit.
- One commit per task. Commit format:
  ```
  feat({{scope}}): {{short description}}

  Refs: .specs/{{feature-slug}}/tasks.md#task-N
  Criterion: {{AC-N text}}
  ```
- Mark each task `DONE` in this file after its commit is pushed
  (via `write_doc`).
- If a task reveals new requirements, stop and update `spec.md`
  before continuing.

---

## Task 1 — {{TASK_TITLE}}

**Status:** TODO | IN PROGRESS | DONE | BLOCKED
**Objective:** {{one sentence — what does completing this task achieve?}}
**Acceptance link:** `spec.md#AC-1`
**Blocked by:** None

**Inputs (files to read):**
- `.specs/{{feature-slug}}/design.md#{{section}}`
- `{{source file or module}}`

**Outputs (files to create or modify):**
- `{{file path}}` — create | modify

**Tests to write:**
- {{test description that directly maps to AC-1}}
- {{edge case from spec.md#edge-cases}}

**Done checklist:**
- [ ] Code written and follows conventions in `_constitution/coding-conventions-index.md`
- [ ] Tests written and passing
- [ ] Criterion AC-1 verifiably satisfied
- [ ] Committed and pushed

---

## Task 2 — {{TASK_TITLE}}

**Status:** TODO
**Objective:** {{one sentence}}
**Acceptance link:** `spec.md#AC-2`
**Blocked by:** Task 1

**Inputs:**
- `{{file path}}`

**Outputs:**
- `{{file path}}` — create | modify

**Tests to write:**
- {{test description}}

**Done checklist:**
- [ ] Code written
- [ ] Tests written and passing
- [ ] Criterion AC-2 satisfied
- [ ] Committed and pushed

---

## Task N — {{TASK_TITLE}}

{{Repeat structure. Every acceptance criterion in spec.md must be covered
by at least one task. Verify coverage before submitting for G3 review.}}

---

## Criterion Coverage Map

{{Fill before G3. Every AC must appear at least once.}}

| Criterion | Covered by task(s) |
|---|---|
| AC-1 | Task 1 |
| AC-2 | Task 2 |
| AC-N | Task N |

---

## Risks

- **{{RISK_1}}:** {{mitigation}}
- **{{RISK_2}}:** {{mitigation}}