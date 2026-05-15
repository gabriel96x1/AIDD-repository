# Evidence: {{Feature Name}}

**Slug:** `{{feature-slug}}`
**Spec:** `.specs/{{feature-slug}}/spec.md`
**Tasks:** `.specs/{{feature-slug}}/tasks.md`
**Status:** OPEN
**Created:** {{DATE}}

---

## Evidence Rules

- Every criterion from `spec.md` must appear here with a status.
- PASS requires a verifiable reference: test file path with line number,
  a CI run link, or a documented manual check with reproduction steps.
- "It works locally" is not evidence.
- PARTIAL is only acceptable with an explanation and a follow-up task reference.
- The agent fills this file as tasks complete (via `write_doc`).
  Humans verify at G4.

---

## Criterion Evidence

### AC-1 — {{Paste criterion text from spec.md}}

**Status:** ✅ PASS | ❌ FAIL | ⚠️ PARTIAL | 🔲 PENDING
**Evidence type:** automated test | CI run | screenshot | log excerpt | manual check
**Reference:**
```
{{file path}}:{{line number}}
— or —
{{CI run URL}}
— or —
Manual check: {{reproduction steps}}
```
**Notes:** {{anything the reviewer needs to interpret the evidence}}

---

### AC-2 — {{Paste criterion text from spec.md}}

**Status:** 🔲 PENDING
**Evidence type:** —
**Reference:** —
**Notes:** —

---

### AC-N — {{Paste criterion text from spec.md}}

{{Repeat for every criterion. Order must match spec.md.}}

---

## Summary

| Criterion | Status |
|---|---|
| AC-1 | 🔲 PENDING |
| AC-2 | 🔲 PENDING |
| AC-N | 🔲 PENDING |

**Overall:** PASS (all criteria) | BLOCKED ({{N}} pending) | FAIL ({{N}} failing)

---

## G4 Checklist

Filled by the human reviewer before closing the gate:

- [ ] All criteria are PASS
- [ ] No criterion is PARTIAL without a follow-up task
- [ ] Test references are reachable (paths exist, CI links resolve)
- [ ] No out-of-scope behavior was introduced (check against spec.md non-goals)
- [ ] `roadmap.md` feature status updated to DONE
- [ ] Replanning triggers checked (see `AGENTS.md`)