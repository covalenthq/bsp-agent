# SECURITY-FIX-PLAN-20260126

## Target Repository
- **Repo**: covalenthq/bsp-agent
- **Branch**: `security/cve-remediation-20260126`
- **Base**: main
- **Ecosystems**: Go

---

## 1. ALERTS SUMMARY

| # | Package | Severity | CVE | CVSS | Summary | Patched Version |
|---|---------|----------|-----|------|---------|-----------------|
| 34 | github.com/ethereum/go-ethereum | HIGH | CVE-2026-22868 | - | DoS via malicious p2p message (high CPU) | 1.16.8 |
| 33 | github.com/ethereum/go-ethereum | HIGH | CVE-2026-22862 | - | DoS via malicious p2p message (crash) | 1.16.8 |
| 32 | golang.org/x/crypto | MEDIUM | CVE-2025-47914 | 5.3 | ssh/agent panic on malformed message | 0.45.0 |
| 31 | golang.org/x/crypto | MEDIUM | CVE-2025-58181 | 5.3 | SSH unbounded memory consumption | 0.45.0 |

**Total**: 4 alerts affecting 2 package(s)

**Current versions in go.mod**:
- `github.com/ethereum/go-ethereum v1.15.11` (direct dependency) → needs update
- `golang.org/x/crypto v0.36.0` (indirect) → needs v0.45.0

**Note**: The go-ethereum alerts reference version 1.16.8 as patched, but current repo uses v1.15.11. This may be a constraint from the ethereum/go-ethereum project versioning. Need to verify the latest available version.

---

## 2. PRE-FLIGHT CHECKS
- [x] Verify clean working directory: `git status`
- [x] Analyze dep graph for golang.org/x/crypto: `go mod graph | grep golang.org/x/crypto`
- [x] Analyze dep graph for go-ethereum: `go mod graph | grep go-ethereum`
- [x] Note: golang.org/x/crypto is indirect - will be updated transitively

**Commit**: None (read-only phase)

---

## 3. BRANCH SETUP
- [x] Fetch latest: `git fetch origin`
- [x] Checkout main: `git checkout main && git pull`
- [x] Create branch: `git checkout -b security/cve-remediation-20260126`
- [x] Push branch: `git push -u origin security/cve-remediation-20260126`

**Commit**: None (branch creation only)

---

## 4. DEPENDENCY UPDATES

### Strategy
- Minimal version bump to exact patched version
- Single commit for all security updates (single PR approach)
- If breaking changes detected → STOP and escalate to human

### Execute Updates

#### 4.1 Update golang.org/x/crypto (indirect → direct bump)
- [x] Update crypto: `go get golang.org/x/crypto@v0.45.0`

#### 4.2 Update go-ethereum
- [x] Check latest version: `go list -m -versions github.com/ethereum/go-ethereum | tail -1`
- [x] Update go-ethereum: `go get github.com/ethereum/go-ethereum@latest`

#### 4.3 Clean up
- [x] `go mod tidy`
- [x] `go mod verify`

**Commit**:
```bash
git add go.mod go.sum
git commit -m "fix(deps): patch CVE-2026-22868, CVE-2026-22862, CVE-2025-47914, CVE-2025-58181

Resolves:
- CVE-2026-22868: go-ethereum DoS via malicious p2p message (high CPU)
- CVE-2026-22862: go-ethereum DoS via malicious p2p message (crash)
- CVE-2025-47914: x/crypto ssh/agent panic on malformed message
- CVE-2025-58181: x/crypto SSH unbounded memory consumption"
```

---

## 5. VERIFICATION LOOP

> **CRITICAL**: Do NOT proceed to PR until ALL checks pass.
> On ANY failure → STOP → analyze if breaking change → escalate if needed.

### 5.1 Static Analysis
- [x] `go vet ./...`
- [x] `go fmt ./...` (verify no changes needed)

**On failure**: If breaking API changes, STOP and escalate.

### 5.2 Build Verification
- [x] `go build ./...`

**On failure**: If compilation errors from dep changes, STOP and escalate.

### 5.3 Test Suite
- [x] `go test ./...`
- [x] `go test -race ./...`

**On failure**: If test failures from dep changes, STOP and escalate.

### 5.4 Security Re-audit
- [x] Run: `govulncheck ./...` (if available)
- [x] Or: check go.sum versions match patched versions
- [x] Confirm patched CVEs no longer appear

### 5.5 Module Verification
- [x] `go mod tidy` (should produce no changes)
- [x] `go mod verify` (should show "all modules verified")

---

## 6. POST-FIX VALIDATION
- [x] All section 5 checks pass
- [x] No new vulnerabilities introduced
- [x] Review diff one final time: `git diff main...HEAD`

**Commit**:
```bash
git add SECURITY-FIX-PLAN.md
git commit -m "chore: mark SECURITY-FIX-PLAN as verified"
```

---

## 7. CREATE PULL REQUEST

```bash
gh pr create \
  --title "fix(security): patch go-ethereum and x/crypto CVEs" \
  --body "$(cat <<'EOF'
## Security Fixes

### Vulnerabilities Addressed

| CVE | Package | Severity | Description |
|-----|---------|----------|-------------|
| CVE-2026-22868 | github.com/ethereum/go-ethereum | HIGH | DoS via malicious p2p message (high CPU) |
| CVE-2026-22862 | github.com/ethereum/go-ethereum | HIGH | DoS via malicious p2p message (crash) |
| CVE-2025-47914 | golang.org/x/crypto | MEDIUM | ssh/agent panic on malformed message |
| CVE-2025-58181 | golang.org/x/crypto | MEDIUM | SSH unbounded memory consumption |

### Changes
- Updated `github.com/ethereum/go-ethereum` to latest patched version
- Updated `golang.org/x/crypto` to v0.45.0

### Verification Completed
- [x] `go vet ./...` - passed
- [x] `go build ./...` - passed
- [x] `go test ./...` - passed
- [x] `go mod verify` - passed

### References
- https://github.com/advisories/GHSA-mq3p-rrmp-79jg
- https://github.com/advisories/GHSA-mr7q-c9w9-wh4h
- https://github.com/advisories/GHSA-f6x5-jh6r-wrfv
- https://github.com/advisories/GHSA-j5w8-q4qc-rx2x

---
Generated by CVE-Manager Agent
EOF
)" \
  --label "security"
```

- [ ] PR created successfully
- [ ] Request review from maintainers

---

## 8. CLEANUP (Post-Merge)
- [ ] Delete SECURITY-FIX-PLAN.md from repo after PR is merged
- [ ] Delete local and remote branch:
      `git push origin --delete security/cve-remediation-20260126`
      `git branch -D security/cve-remediation-20260126`

---

## 9. ESCALATION PROTOCOL

If ANY of these occur, STOP and escalate:
- Compilation errors after dependency update
- Test failures after dependency update
- Breaking API changes requiring code modifications
- Dependency conflicts preventing update
- Transitive dependency requiring major version bump of direct dep

**To escalate** (update plan.md in cve-manager root):

1. Change repo line from:
   ```
   - [ ] covalenthq/bsp-agent
   ```
   To:
   ```
   - [!] covalenthq/bsp-agent **BLOCKED** - {reason}
   ```

2. Add details under `## Escalations` section in plan.md

---

## EXECUTION LOG

| Step | Status | Timestamp | Notes |
|------|--------|-----------|-------|
| 2. Pre-flight | PASS | 2026-01-29 16:03:20 | go toolchain available; dependency graph reviewed |
| 3. Branch setup | PASS | 2026-01-29 16:05:20 | security/cve-remediation-20260126 created and pushed |
| 4. Dep updates | PASS | 2026-01-29 16:09:30 | Updated go-ethereum to v1.16.8 and x/crypto to v0.45.0 |
| 5. Verification | PASS | 2026-01-29 16:16:10 | govulncheck not available; verified go.mod versions and tests/build pass |
| 6. Validation | PASS | 2026-01-29 16:19:05 | Verification complete; diff reviewed |
| 7. PR created | ⬜ | | |
| 8. Cleanup | ⬜ | | |

---

## AGENT INSTRUCTIONS

1. Execute sections 2-7 sequentially
2. Mark checkboxes [x] as you complete each step
3. Commit after each major section (as indicated)
4. On verification failure: analyze error → if breaking change → STOP and escalate
5. Do NOT create PR until section 5 fully passes
6. Update EXECUTION LOG with timestamps and notes
7. After PR merge confirmation, execute section 8 cleanup
