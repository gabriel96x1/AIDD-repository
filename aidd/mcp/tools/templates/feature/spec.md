# Spec: {{Feature Name}}

**Slug:** `{{feature-slug}}`
**Status:** DRAFT
**Created:** {{DATE}}
**Author:** {{agent | human name}}
**Linked roadmap phase:** Phase {{N}} — {{phase name}}

---

## Goal

{{One paragraph. What does this change achieve and why does it matter to the
project's mission? If it doesn't map to the mission, it shouldn't be built.}}

---

## Non-Goals

{{Hard boundary. The agent must not implement anything on this list, even if it
seems obviously useful. Be explicit — vague non-goals create scope creep.}}

- This feature does not {{NON_GOAL_1}}.
- This feature does not {{NON_GOAL_2}}.
- This feature does not {{NON_GOAL_3}}.

---

## User Stories

- As a {{ROLE}}, I want {{CAPABILITY}}, so that {{OUTCOME}}.
- As a {{ROLE}}, I want {{CAPABILITY}}, so that {{OUTCOME}}.

---

## Acceptance Criteria

{{EARS notation only. One pattern per criterion. Must be testable.
Use "shall" for mandatory behavior. Never "should", "might", "could".
Number every criterion — tasks and evidence reference them by number.}}

**AC-1.** WHEN {{trigger}} THE system SHALL {{response}}.

**AC-2.** IF {{condition}} THEN THE system SHALL {{response}}.

**AC-3.** WHILE {{state}} THE system SHALL {{behavior}}.

**AC-4.** The system shall {{always-true behavior}}.

**AC-5.** WHERE {{optional feature is enabled}} THE system SHALL {{behavior}}.

---

## Edge Cases

- **{{EDGE_CASE_1}}:** {{expected behavior}}
- **{{EDGE_CASE_2}}:** {{expected behavior}}
- **{{EDGE_CASE_3}}:** {{expected behavior}}

---

## Open Questions

| # | Question | Owner | Resolved | Resolution |
|---|---|---|---|---|
| 1 | {{QUESTION}} | {{human/agent}} | No | — |

---

## Out of Scope (adjacent features)

- {{ADJACENT_FEATURE_1}} — deferred to Phase {{N}}.
- {{ADJACENT_FEATURE_2}} — separate spec: `{{other-slug}}`.

---

## References

- Mission alignment: `.specs/_constitution/mission.md`
- Related spec: `.specs/{{related-slug}}/spec.md`
- Tech decision: `.specs/_constitution/tech-stack.md#{{decision}}`