# Design: {{Feature Name}}

**Slug:** `{{feature-slug}}`
**Spec:** `.specs/{{feature-slug}}/spec.md`
**Status:** DRAFT
**Created:** {{DATE}}
**Author:** {{agent | human name}}

---

## Architecture Overview

{{Which components, layers, or services are touched by this change.
Use a diagram (ASCII or Mermaid) when the change spans more than two layers.
Be specific about what is new vs what is modified.}}

```
{{diagram or description}}
```

**Layers touched:**
- [ ] {{LAYER_1}} — new | modified | read-only
- [ ] {{LAYER_2}} — new | modified | read-only

---

## Data Model

{{New or modified schemas, types, interfaces, or database tables.
Include field names, types, constraints, and migration notes if applicable.}}

```
{{schema or type definition}}
```

**Migration required:** Yes | No
**Migration notes:** {{If yes, describe the migration strategy and rollback plan.}}

---

## API Contracts

{{Input/output signatures for every new or modified interface.
Be explicit about types — no "object" or "any".}}

### {{CONTRACT_NAME}}

**Input:**
```
{{type definition or schema}}
```

**Output:**
```
{{type definition or schema}}
```

**Error cases:**
- `{{ERROR_CODE}}` — {{when this is returned}}

---

## New Dependencies

{{Every new runtime dependency introduced by this feature.
If none, write "None." Each entry must justify why existing options are insufficient.}}

| Package | Version | Justification | Existing alternative considered |
|---|---|---|---|
| `{{package}}` | `{{version}}` | {{why needed}} | {{what was evaluated and why it didn't fit}} |

---

## Rejected Alternatives

{{Mandatory. Prevents the same debate from recurring and gives future agents
context for why the current design exists.}}

### Alternative: {{ALTERNATIVE_1}}
**Why rejected:** {{reason}}

### Alternative: {{ALTERNATIVE_2}}
**Why rejected:** {{reason}}

---

## Rollout Plan

**Feature flag required:** Yes | No
**Flag name:** `{{FLAG_NAME}}` (if applicable)
**Backward compatible:** Yes | No
**Rollback procedure:** {{how to revert if something goes wrong post-deploy}}
**Monitoring:** {{what metrics or logs indicate the feature is working correctly}}

---

## Security & Privacy Considerations

{{Does this change touch auth, user data, external APIs, or the filesystem?
If yes, describe how security and privacy are handled. If no, write "None."}}

---

## Open Questions

| # | Question | Owner | Resolved | Resolution |
|---|---|---|---|---|
| 1 | {{QUESTION}} | {{human/agent}} | No | — |