# Go SDK Structural Parity with TypeScript SDK

**Date:** 2026-02-28
**Status:** Approved

---

## Goal

Bring the Go SDK (`github.com/maximiseai/enrich-go-sdk`) to the same structural level as the TypeScript SDK by adding CI/CD workflows, documentation, licensing, and tooling config that currently only exist in the TypeScript repo.

---

## Files to Add

| File | TypeScript Equivalent | Notes |
|---|---|---|
| `.github/workflows/ci.yml` | same | Go: `go test ./...`, `go vet`, `go build ./...` |
| `.github/workflows/release.yml` | same | Uses GoReleaser instead of semantic-release |
| `.goreleaser.yml` | `.releaserc.json` | GoReleaser config for library-only releases |
| `.commitlintrc.json` | same | Identical conventional commits config |
| `README.md` | same | Full Go SDK docs with Go syntax examples |
| `CHANGELOG.md` | same | Initial changelog |
| `LICENSE` | same | MIT license |
| `.gitignore` | same | Go-specific ignores |

## Files Not Applicable in Go

- `.husky/` — Node.js only; commit linting enforced in CI instead
- `.npmignore` / `.npmrc` — npm-specific, no Go equivalent needed

---

## go.mod Module Path Update

Change `module sdk` → `module github.com/maximiseai/enrich-go-sdk`.

This requires a global find+replace of `"sdk/` → `"github.com/maximiseai/enrich-go-sdk/` across all `.go` files.

---

## CI Workflow

Triggers on `push` and `pull_request` to `main`. Steps:
1. Checkout
2. Setup Go (go version from go.mod: 1.21)
3. `go vet ./...`
4. `go test ./...`
5. `go build ./...`

## Release Workflow

Triggers on `push` to `main` (same as TypeScript). Steps:
1. Checkout with full history (`fetch-depth: 0`)
2. Setup Go
3. Run GoReleaser (reads conventional commits → determines semver bump → creates git tag → creates GitHub Release with changelog)

Skip condition: `!contains(github.event.head_commit.message, 'chore(release)')` (same as TypeScript)

## GoReleaser Config

Configured as a **library** (no binary builds):
- `builds: []` — skip binary compilation
- `changelog.use: conventional-commits` — parse commit history for release notes
- `release.github: true` — publish to GitHub Releases
- Auto-generates CHANGELOG entries

## Commit Linting

Uses the same `.commitlintrc.json` as TypeScript (`@commitlint/config-conventional`). Since Husky is Node.js-only, enforcement happens in CI via a `commitlint` GitHub Actions step rather than a local git hook.

---

## Required GitHub Secrets

| Secret | Purpose |
|---|---|
| `GITHUB_TOKEN` | Auto-provided by GitHub Actions; used by GoReleaser to create releases |

No `NPM_TOKEN` needed (Go modules are published via git tags, not a registry).
