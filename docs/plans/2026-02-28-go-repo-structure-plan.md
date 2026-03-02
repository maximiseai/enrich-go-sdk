# Go SDK Structural Parity Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add CI/CD workflows, README, LICENSE, CHANGELOG, .gitignore, and tooling config to the Go SDK so it matches the TypeScript SDK's structure.

**Architecture:** Update the go.mod module path to `github.com/maximiseai/enrich-go-sdk`, fix all internal imports via find+replace, then add each missing file — no existing logic is changed.

**Tech Stack:** Go 1.21+, GitHub Actions, GoReleaser v2, mathieudutour/github-tag-action for conventional-commit semver tagging.

---

### Task 1: Update go.mod module path

**Files:**
- Modify: `go/go.mod:1`

**Step 1: Edit go.mod**

Open `go/go.mod` and change line 1 from:
```
module sdk
```
to:
```
module github.com/maximiseai/enrich-go-sdk
```

Full resulting go.mod (do not change any other line):
```
module github.com/maximiseai/enrich-go-sdk

go 1.21

toolchain go1.23.8

require github.com/google/uuid v1.6.0

require github.com/stretchr/testify v1.8.4

require gopkg.in/yaml.v3 v3.0.1 // indirect

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)
```

**Step 2: Verify**

Run: `head -1 go/go.mod`
Expected output: `module github.com/maximiseai/enrich-go-sdk`

**Step 3: Commit**

```bash
cd go
git add go.mod
git commit -m "chore: update module path to github.com/maximiseai/enrich-go-sdk"
```

---

### Task 2: Update all internal imports

**Files:**
- Modify: all `*.go` files (bulk find+replace — do NOT edit manually)

Every `.go` file uses `"sdk/..."` or `sdk "sdk"` as import paths. These must all be updated to use the new module path.

**Step 1: Replace all sub-package imports**

Run from the `go/` directory:
```bash
cd go
find . -name "*.go" -exec sed -i 's|"sdk/|"github.com/maximiseai/enrich-go-sdk/|g' {} \;
```

**Step 2: Replace root package aliased imports**

```bash
find . -name "*.go" -exec sed -i 's|sdk "sdk"|sdk "github.com/maximiseai/enrich-go-sdk"|g' {} \;
```

**Step 3: Verify a sample file**

Run: `head -20 client/client.go`

Expected — imports should now read:
```go
import (
	companyfollowers "github.com/maximiseai/enrich-go-sdk/companyfollowers"
	core "github.com/maximiseai/enrich-go-sdk/core"
	emailfinder "github.com/maximiseai/enrich-go-sdk/emailfinder"
	...
)
```

**Step 4: Verify the build compiles**

```bash
go build ./...
```
Expected: no output (success). Fix any remaining import errors before continuing.

**Step 5: Run tests**

```bash
go test ./... -count=1 -short
```
Expected: all tests pass (or skip if they require external services).

**Step 6: Commit**

```bash
git add -A
git commit -m "chore: update all internal imports to new module path"
```

---

### Task 3: Add .gitignore

**Files:**
- Create: `go/.gitignore`

**Step 1: Create the file**

Create `go/.gitignore` with this content:
```gitignore
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of go coverage tool
*.out
coverage.html

# Go build cache (local, not needed in repo)
/dist/

# GoReleaser artifacts
dist/

# Environment variables
.env
.env.*
!.env.example

# OS files
.DS_Store
Thumbs.db

# IDE
.vscode/
.idea/
*.swp
*.swo

# Logs
*.log
```

**Step 2: Commit**

```bash
cd go
git add .gitignore
git commit -m "chore: add .gitignore"
```

---

### Task 4: Add LICENSE

**Files:**
- Create: `go/LICENSE`

**Step 1: Create the file**

Create `go/LICENSE` with this exact content (update year to 2025):
```
MIT License

Copyright (c) 2025 Enrich

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

**Step 2: Commit**

```bash
cd go
git add LICENSE
git commit -m "chore: add MIT license"
```

---

### Task 5: Add CHANGELOG.md

**Files:**
- Create: `go/CHANGELOG.md`

**Step 1: Create the file**

Create `go/CHANGELOG.md`:
```markdown
# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2026-02-28

### Added

- Initial release of the Enrich API Go SDK.
- Email Finder: single and batch email lookups.
- Email Validation: single and batch email validation.
- Phone Finder: single and batch phone lookups.
- Reverse Email Lookup: single and bulk reverse lookups.
- Lead Finder: search, count, reveal, enrich, export, saved searches.
- People Search: employee finder and waterfall ICP search.
- Company Followers: follower jobs, count estimates, CSV export.
- Teams: member management and invitations.
- Wallets: balance and transaction history.
- Automatic retries with configurable max attempts.
- Configurable HTTP client and base URL.
- Raw response access via `WithRawResponse` client.
- Typed error types for all HTTP status codes.
```

**Step 2: Commit**

```bash
cd go
git add CHANGELOG.md
git commit -m "chore: add CHANGELOG.md"
```

---

### Task 6: Add .commitlintrc.json

**Files:**
- Create: `go/.commitlintrc.json`

**Step 1: Create the file**

Create `go/.commitlintrc.json` (identical to the TypeScript repo):
```json
{
  "extends": ["@commitlint/config-conventional"]
}
```

**Step 2: Commit**

```bash
cd go
git add .commitlintrc.json
git commit -m "chore: add commitlint config for conventional commits"
```

---

### Task 7: Add .goreleaser.yml

**Files:**
- Create: `go/.goreleaser.yml`

This is the Go equivalent of `.releaserc.json`. It tells GoReleaser how to create GitHub releases for a library (no binary builds).

**Step 1: Create the file**

Create `go/.goreleaser.yml`:
```yaml
version: 2

project_name: enrich-go-sdk

before:
  hooks:
    - go mod tidy

# This is a library — skip binary compilation
builds:
  - skip: true

changelog:
  use: conventional-commits
  sort: desc
  abbrev: -1
  groups:
    - title: "Features"
      regexp: "^feat"
      order: 0
    - title: "Bug Fixes"
      regexp: "^fix"
      order: 1
    - title: "Documentation"
      regexp: "^docs"
      order: 2
    - title: Other
      order: 999
  filters:
    exclude:
      - "^chore(release)"
      - "^ci"
      - "^test"

release:
  github:
    owner: maximiseai
    name: enrich-go-sdk
  draft: false
  prerelease: auto
  name_template: "v{{.Version}}"
```

**Step 2: Commit**

```bash
cd go
git add .goreleaser.yml
git commit -m "chore: add goreleaser config for library releases"
```

---

### Task 8: Add GitHub Actions CI workflow

**Files:**
- Create: `go/.github/workflows/ci.yml`

**Step 1: Create directory**

```bash
mkdir -p go/.github/workflows
```

**Step 2: Create the file**

Create `go/.github/workflows/ci.yml`:
```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: true

      - name: Vet
        run: go vet ./...

      - name: Test
        run: go test ./... -count=1 -short

      - name: Build
        run: go build ./...
```

**Step 3: Commit**

```bash
cd go
git add .github/workflows/ci.yml
git commit -m "ci: add CI workflow for Go"
```

---

### Task 9: Add GitHub Actions release workflow

**Files:**
- Create: `go/.github/workflows/release.yml`

This mirrors the TypeScript release workflow. It uses `mathieudutour/github-tag-action` to read conventional commits and create the semver git tag, then runs GoReleaser to create the GitHub release.

**Step 1: Create the file**

Create `go/.github/workflows/release.yml`:
```yaml
name: Release

on:
  push:
    branches: [main]

permissions:
  contents: write
  issues: write
  pull-requests: write

jobs:
  release:
    runs-on: ubuntu-latest
    # Skip release commits made by the tag action
    if: "!contains(github.event.head_commit.message, 'chore(release)')"

    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
          persist-credentials: true

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: true

      - name: Build
        run: go build ./...

      - name: Bump version and create tag
        id: tag
        uses: mathieudutour/github-tag-action@v6.2
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          default_bump: patch
          tag_prefix: "v"

      - name: Release with GoReleaser
        uses: goreleaser/goreleaser-action@v6
        if: steps.tag.outputs.new_tag != ''
        with:
          version: latest
          args: release --clean --skip=build
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**Step 2: Commit**

```bash
cd go
git add .github/workflows/release.yml
git commit -m "ci: add release workflow with GoReleaser"
```

---

### Task 10: Add README.md

**Files:**
- Create: `go/README.md`

**Step 1: Create the file**

Create `go/README.md` with the full Go-syntax documentation:

```markdown
# Enrich API Go SDK

The official Go SDK for the [Enrich API](https://enrich.so) — find emails, phone numbers, validate emails, discover leads, and more.

## Installation

```bash
go get github.com/maximiseai/enrich-go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	"github.com/maximiseai/enrich-go-sdk/client"
	"github.com/maximiseai/enrich-go-sdk/option"
	sdk "github.com/maximiseai/enrich-go-sdk"
)

func main() {
	c := client.NewClient(
		option.WithAPIKey("YOUR_API_KEY"),
		option.WithBaseURL("https://api.enrich.so/v3"),
	)

	result, err := c.EmailFinder.FindEmail(context.Background(), &sdk.EmailFinderRequest{
		FirstName: "Emily",
		LastName:  "Zhang",
		Domain:    "figma.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Data.Email)
}
```

## Authentication

All API requests require an API key passed via the `x-api-key` header. Provide it when creating the client:

```go
c := client.NewClient(
	option.WithAPIKey("YOUR_API_KEY"),
	option.WithBaseURL("https://api.enrich.so/v3"),
)
```

You can also read it from an environment variable:

```go
import "os"

c := client.NewClient(
	option.WithAPIKey(os.Getenv("ENRICH_API_KEY")),
	option.WithBaseURL("https://api.enrich.so/v3"),
)
```

## Credit Costs

Every API call costs credits. You are **not charged** when a lookup fails (e.g. email not found).

| API | Method | Cost |
|---|---|---|
| Email Finder | `FindEmail` | 10 credits |
| Email Finder | `BatchFindEmails` | 10 credits/lead |
| Email Finder | `GetEmailFinderBatchStatus` | Free |
| Email Finder | `GetEmailFinderBatchResults` | Free |
| Email Validation | `ValidateEmail` | 1 credit |
| Email Validation | `BatchValidateEmails` | 1 credit/email |
| Email Validation | `GetEmailValidationBatchStatus` | Free |
| Email Validation | `GetEmailValidationBatchResults` | Free |
| Phone Finder | `PhoneLookup` | 500 credits |
| Phone Finder | `CreatePhoneBulkJob` | 10 credits/lead |
| Phone Finder | `GetPhoneBulkJobStatus` | Free |
| Phone Finder | `GetPhoneBulkJobResults` | Free |
| Reverse Email Lookup | `ReverseEmailLookup` | 10 credits |
| Reverse Email Lookup | `ReverseEmailBulk` | 10 credits/email |
| Reverse Email Lookup | `GetBulkLookupStatus` | Free |
| Reverse Email Lookup | `GetBulkLookupResults` | Free |
| Lead Finder | `Search` | Free (preview) |
| Lead Finder | `Count` | Free |
| Lead Finder | `Reveal` | Credits per lead |
| Lead Finder | `Enrich` | Credits per lead |
| Lead Finder | `UnlockNames` | Credits per lead |
| Lead Finder | `Export` | Credits per lead |
| Lead Finder | `GetFilterOptions` | Free |
| Lead Finder | `SuggestCompany` | Free |
| Lead Finder | `SavedSearches` | Free |
| People Search | `EmployeeFinder` | Credits per result |
| People Search | `WaterfallIcpSearch` | Credits per result |
| Company Followers | `StartScrape` | Credits per follower |
| Company Followers | `StartCountEstimate` | 100 credits |
| Company Followers | `CheckLimit` | Free |
| Company Followers | All status/results | Free |
| Teams | All methods | Free |
| Wallets | All methods | Free |

Check your balance anytime with `c.Wallets.Balance(ctx)`.

## API Reference

### Email Finder

Find professional email addresses given a person's name and company domain.

```go
// Single lookup (10 credits, free if not found)
result, err := c.EmailFinder.FindEmail(ctx, &sdk.EmailFinderRequest{
	FirstName: "Emily",
	LastName:  "Zhang",
	Domain:    "figma.com",
})

// Batch lookup
webhookURL := "https://api.yourapp.com/webhooks/enrich"
batch, err := c.EmailFinder.BatchFindEmails(ctx, &sdk.BatchEmailFinderRequest{
	Leads: []*sdk.Lead{
		{FirstName: "Emily", LastName: "Zhang", Domain: "figma.com"},
		{FirstName: "David", LastName: "Kim", Domain: "vercel.com"},
	},
	WebhookURL: &webhookURL,
})

// Check batch status
status, err := c.EmailFinder.GetEmailFinderBatchStatus(ctx, &sdk.GetEmailFinderBatchStatusRequest{
	BatchID: batch.Data.BatchID,
})

// Get batch results (paginated)
page := 1
limit := 50
results, err := c.EmailFinder.GetEmailFinderBatchResults(ctx, &sdk.GetEmailFinderBatchResultsRequest{
	BatchID: batch.Data.BatchID,
	Page:    &page,
	Limit:   &limit,
})
```

### Email Validation

Validate email addresses for deliverability.

```go
// Single validation
validation, err := c.EmailValidation.ValidateEmail(ctx, &sdk.EmailValidationRequest{
	Email: "emily@figma.com",
})

// Batch validation
batch, err := c.EmailValidation.BatchValidateEmails(ctx, &sdk.BatchEmailValidationRequest{
	Emails: []string{"emily@figma.com", "david@vercel.com"},
})

// Check status
status, err := c.EmailValidation.GetEmailValidationBatchStatus(ctx, &sdk.GetEmailValidationBatchStatusRequest{
	BatchID: batch.Data.BatchID,
})

// Get results
results, err := c.EmailValidation.GetEmailValidationBatchResults(ctx, &sdk.GetEmailValidationBatchResultsRequest{
	BatchID: batch.Data.BatchID,
})
```

### Phone Finder

Find phone numbers for professionals.

```go
// Single lookup
linkedin := "https://linkedin.com/in/emilyzhang"
phone, err := c.PhoneFinder.PhoneLookup(ctx, &sdk.PhoneLookupRequest{
	Linkedin: &linkedin,
})

// Batch lookup
batch, err := c.PhoneFinder.CreatePhoneBulkJob(ctx, &sdk.PhoneBatchRequest{
	Linkedins: []string{"https://linkedin.com/in/emilyzhang"},
})

// Check status & get results
status, err := c.PhoneFinder.GetPhoneBulkJobStatus(ctx, &sdk.GetPhoneBulkJobStatusRequest{
	JobID: batch.Data.BatchID,
})

results, err := c.PhoneFinder.GetPhoneBulkJobResults(ctx, &sdk.GetPhoneBulkJobResultsRequest{
	JobID: batch.Data.BatchID,
})
```

### Reverse Email Lookup

Look up a person's LinkedIn profile using their email address.

```go
// Single lookup
profile, err := c.ReverseEmailLookup.ReverseEmailLookup(ctx, &sdk.LookupRequest{
	Email: "emily@figma.com",
})

// Bulk lookup
bulk, err := c.ReverseEmailLookup.ReverseEmailBulk(ctx, &sdk.BulkLookupRequest{
	Emails: []string{"emily@figma.com", "david@vercel.com"},
})

// Check status & get results
status, err := c.ReverseEmailLookup.GetBulkLookupStatus(ctx, &sdk.GetBulkLookupStatusRequest{
	BatchID: bulk.Data.BatchID,
})

results, err := c.ReverseEmailLookup.GetBulkLookupResults(ctx, &sdk.GetBulkLookupResultsRequest{
	BatchID: bulk.Data.BatchID,
})
```

### Lead Finder

Search and discover leads based on filters.

```go
// Search leads
leads, err := c.LeadFinder.Search(ctx, &sdk.LeadSearchRequest{
	JobTitle: []string{"CTO"},
})

// Get lead count
count, err := c.LeadFinder.Count(ctx, &sdk.LeadCountRequest{
	JobTitle: []string{"CTO"},
})

// Reveal lead contact details
revealed, err := c.LeadFinder.Reveal(ctx, &sdk.LeadRevealRequest{
	Leads: []*sdk.LeadRevealItem{
		{Id: "lead_123"},
	},
})

// Enrich leads
enriched, err := c.LeadFinder.Enrich(ctx, &sdk.LeadEnrichRequest{
	Leads:  []*sdk.LeadRevealItem{{Id: "lead_123"}},
	Fields: []string{"email"},
})

// List saved searches
searches, err := c.LeadFinder.SavedSearches(ctx)

// Suggest company names
suggestions, err := c.LeadFinder.SuggestCompany(ctx, &sdk.SuggestCompanyNamesRequest{
	Q: "Goo",
})

// Get filter options
options, err := c.LeadFinder.GetFilterOptions(ctx)

// Unlock names
unlocked, err := c.LeadFinder.UnlockNames(ctx, &sdk.UnlockNamesRequest{
	Leads: []*sdk.LeadRevealItem{{Id: "lead_123"}},
})
```

### People Search

Find employees at specific companies.

```go
// Find employees
employees, err := c.PeopleSearch.EmployeeFinder(ctx, &sdk.EmployeeFinderRequest{
	CompanyLinkedinUrl: "https://linkedin.com/company/figma",
})

// Waterfall ICP search
icpResults, err := c.PeopleSearch.WaterfallIcpSearch(ctx, &sdk.WaterfallIcpSearchRequest{
	CompanyLinkedinUrl: "https://linkedin.com/company/figma",
})
```

### Company Followers

Get followers of a LinkedIn company page.

```go
maxLimit := 1000

// Start a follower scrape job
job, err := c.CompanyFollowers.StartScrape(ctx, &sdk.CompanyFollowerRequest{
	CompanyUrl: "https://linkedin.com/company/figma",
	MaxLimit:   maxLimit,
})

// Check progress
progress, err := c.CompanyFollowers.ScrapeProgress(ctx, &sdk.GetCompanyFollowerProgressRequest{
	BatchID: job.Data.BatchID,
})

// Get results (paginated)
page := 1
limit := 50
followers, err := c.CompanyFollowers.ScrapeResults(ctx, &sdk.GetCompanyFollowerResultsRequest{
	BatchID: job.Data.BatchID,
	Page:    &page,
	Limit:   &limit,
})

// Start count estimate (100 credits)
estimate, err := c.CompanyFollowers.StartCountEstimate(ctx, &sdk.CountEstimateRequest{
	CompanyUrl: "https://linkedin.com/company/figma",
})

// Check estimate status
estimateStatus, err := c.CompanyFollowers.GetCountEstimateStatus(ctx, &sdk.GetCountEstimateStatusRequest{
	BatchID: estimate.Data.BatchID,
})

// Check daily limit (free)
limit_, err := c.CompanyFollowers.CheckLimit(ctx)
```

### Teams

Manage team members and invitations.

```go
// List members
members, err := c.Teams.Members(ctx, &sdk.ListTeamMembersRequest{
	TeamID: "team_123",
})

// Invite a member
role := "member"
invitation, err := c.Teams.Invite(ctx, &sdk.InviteBody{
	TeamID: "team_123",
	Email:  "newmember@company.com",
	Role:   &role,
})

// List invitations
invitations, err := c.Teams.Invitations(ctx, &sdk.ListTeamInvitationsRequest{
	TeamID: "team_123",
})

// Cancel invitation
err = c.Teams.CancelInvitation(ctx, &sdk.CancelInvitationRequest{
	TeamID:       "team_123",
	InvitationID: "inv_123",
})
```

### Wallets

Check credit balance and transaction history.

```go
// Get wallet balance
balance, err := c.Wallets.Balance(ctx)

// Get transaction history
page := 1
limit := 50
transactions, err := c.Wallets.Transactions(ctx, &sdk.GetWalletTransactionsRequest{
	Page:  &page,
	Limit: &limit,
})
```

## Configuration

### Custom Base URL

```go
c := client.NewClient(
	option.WithAPIKey("YOUR_API_KEY"),
	option.WithBaseURL("https://api.enrich.so/v3"),
)
```

### Retries

Requests are retried automatically. Configure the max number of attempts:

```go
// Client-wide (default is 2)
c := client.NewClient(
	option.WithAPIKey("YOUR_API_KEY"),
	option.WithBaseURL("https://api.enrich.so/v3"),
	option.WithMaxAttempts(5),
)

// Per-request (overrides client default)
result, err := c.EmailFinder.FindEmail(
	ctx,
	&sdk.EmailFinderRequest{FirstName: "Emily", LastName: "Zhang", Domain: "figma.com"},
	option.WithMaxAttempts(0), // Disable retries for this call
)
```

### Custom HTTP Client

```go
import "net/http"
import "time"

httpClient := &http.Client{Timeout: 30 * time.Second}

c := client.NewClient(
	option.WithAPIKey("YOUR_API_KEY"),
	option.WithBaseURL("https://api.enrich.so/v3"),
	option.WithHTTPClient(httpClient),
)
```

### Raw Response Access

Every client has a `WithRawResponse` variant that returns status code, headers, and body:

```go
response, err := c.EmailFinder.WithRawResponse.FindEmail(ctx, &sdk.EmailFinderRequest{
	FirstName: "Emily",
	LastName:  "Zhang",
	Domain:    "figma.com",
})
if err != nil {
	return
}

fmt.Println(response.StatusCode)  // 200
fmt.Println(response.Header)      // http.Header
fmt.Println(response.Body.Data.Email)
```

## Error Handling

The SDK returns typed errors for different HTTP status codes. Use `errors.As` to match them:

```go
import "errors"

result, err := c.EmailFinder.FindEmail(ctx, &sdk.EmailFinderRequest{
	FirstName: "Emily",
	LastName:  "Zhang",
	Domain:    "figma.com",
})
if err != nil {
	var badReq *sdk.BadRequestError
	var unauthorized *sdk.UnauthorizedError
	var paymentRequired *sdk.PaymentRequiredError
	var forbidden *sdk.ForbiddenError
	var notFound *sdk.NotFoundError
	var tooMany *sdk.TooManyRequestsError
	var serverErr *sdk.InternalServerError

	switch {
	case errors.As(err, &badReq):
		// 400 — invalid request parameters
		fmt.Println("Bad request:", badReq.Body)
	case errors.As(err, &unauthorized):
		// 401 — invalid or missing API key
		fmt.Println("Unauthorized:", unauthorized.Body)
	case errors.As(err, &paymentRequired):
		// 402 — insufficient credits
		fmt.Println("Insufficient credits:", paymentRequired.Body)
	case errors.As(err, &forbidden):
		// 403 — not approved for this resource
		fmt.Println("Forbidden:", forbidden.Body)
	case errors.As(err, &notFound):
		// 404 — resource not found
		fmt.Println("Not found:", notFound.Body)
	case errors.As(err, &tooMany):
		// 429 — rate limit exceeded
		fmt.Println("Rate limited:", tooMany.Body)
	case errors.As(err, &serverErr):
		// 500 — server error
		fmt.Println("Server error:", serverErr.Body)
	default:
		fmt.Println("Unexpected error:", err)
	}
	return
}
```

## License

MIT
```

**Step 2: Verify Markdown renders cleanly** — open in a Markdown preview and check all code blocks close properly.

**Step 3: Commit**

```bash
cd go
git add README.md
git commit -m "docs: add comprehensive Go SDK README"
```

---

### Task 11: Final verification

**Step 1: Check all expected files exist**

```bash
cd go
ls -la .gitignore LICENSE CHANGELOG.md .commitlintrc.json .goreleaser.yml README.md
ls -la .github/workflows/ci.yml .github/workflows/release.yml
```

Expected: all 8 files listed with sizes > 0.

**Step 2: Verify no stale `sdk/` imports remain**

```bash
grep -r '"sdk/' . --include="*.go" | grep -v "_test.go"
```

Expected: no output (all replaced).

**Step 3: Final build + test**

```bash
go build ./...
go vet ./...
```

Expected: no errors.

**Step 4: Summary commit (if any loose files)**

```bash
git status
```

If anything unstaged, add and commit with:
```bash
git add -A
git commit -m "chore: finalize go repo structure parity with typescript sdk"
```

---

## Required GitHub Secrets

After pushing to the remote repository, add this secret in **GitHub → Settings → Secrets → Actions**:

| Secret | Purpose | Value |
|---|---|---|
| `GITHUB_TOKEN` | Auto-provided by GitHub Actions — no setup needed | (automatic) |

No additional secrets needed. GoReleaser uses `GITHUB_TOKEN` (auto-provided) to create releases. No npm token required.
