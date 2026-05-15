# Tech Stack — {{PROJECT_NAME}}

**Last updated:** {{DATE}}

All decisions below are active constraints. The agent must not deviate from
them without first adding a new entry here and getting human approval at G2.

---

## Decision Format

### Decision: {{TITLE}}
**Choice:** {{what was chosen}}
**Rationale:** {{why this over alternatives}}
**Constraints:**
- {{constraint}}
**Review trigger:** {{what would cause revisiting this decision}}
**Status:** ACTIVE | SUPERSEDED | UNDER_REVIEW

---

## Language & Runtime

### Decision: Primary language
**Choice:** {{LANGUAGE}} ({{VERSION}})
**Rationale:** {{why}}
**Constraints:**
- All source files must use this language.
- No polyglot mixing without a new decision entry.
**Review trigger:** Major version EOL or team-wide skill shift.
**Status:** ACTIVE

---

## Frameworks & Core Libraries

### Decision: {{FRAMEWORK_NAME}}
**Choice:** {{FRAMEWORK}} ({{VERSION}})
**Rationale:** {{why}}
**Constraints:**
- {{constraint 1}}
- {{constraint 2}}
**Review trigger:** {{trigger}}
**Status:** ACTIVE

---

## State Management

### Decision: {{STATE_SOLUTION}}
**Choice:** {{LIBRARY}} ({{VERSION}})
**Rationale:** {{why}}
**Constraints:**
- {{constraint}}
**Review trigger:** {{trigger}}
**Status:** ACTIVE

---

## Data Persistence

### Decision: {{PERSISTENCE_SOLUTION}}
**Choice:** {{TECHNOLOGY}}
**Rationale:** {{why}}
**Constraints:**
- {{constraint}}
**Review trigger:** {{trigger}}
**Status:** ACTIVE

---

## Testing

### Decision: Testing stack
**Choice:** {{TEST_FRAMEWORK}} for unit/integration, {{E2E_FRAMEWORK}} for E2E.
**Rationale:** {{why}}
**Constraints:**
- Minimum {{N}}% coverage on `{{CORE_DIRS}}`.
- Every acceptance criterion in a spec must have a corresponding automated test.
- No merging without passing test suite.
**Review trigger:** Framework abandonment or better DX alternative.
**Status:** ACTIVE

---

## CI/CD & Deployment

### Decision: {{CI_PLATFORM}}
**Choice:** {{PLATFORM}} — {{brief description}}
**Rationale:** {{why}}
**Constraints:**
- No manual deploys to production. All changes go through pipeline.
- Feature branches require passing CI before PR can be reviewed.
**Review trigger:** Cost, vendor lock-in concerns.
**Status:** ACTIVE

---

## Dependency Policy

- New runtime dependencies require a `design.md` entry at G2 before being added.
- Prefer existing approved libraries over new ones.
- Security advisories trigger immediate review regardless of gate position.
- Dependencies are pinned to exact versions in lockfile.

---

## Update Policy

Add new entries at the bottom of each section. Never delete superseded
decisions — mark them SUPERSEDED and reference the new decision that replaces
them. This preserves the full decision trail.