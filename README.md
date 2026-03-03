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
| Lead Finder | `RevealLeads` | Credits per lead |
| Lead Finder | `EnrichLeads` | Credits per lead |
| Lead Finder | `GetRevealJob` | Free |
| Lead Finder | `ListRevealJobs` | Free |
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

// Submit a reveal job (async — returns a job ticket immediately)
job, err := c.LeadFinder.RevealLeads(ctx, &sdk.LeadRevealRequest{
	Leads: []*sdk.LeadRevealItem{{Id: "lead_123"}},
})
jobID := job.Data.JobID

// Submit an enrich job (async — returns a job ticket immediately)
job, err = c.LeadFinder.EnrichLeads(ctx, &sdk.LeadEnrichRequest{
	Leads:  []*sdk.LeadRevealItem{{Id: "lead_123"}},
	Fields: []string{"email"},
})

// Poll a reveal/enrich job
poll, err := c.LeadFinder.GetRevealJob(ctx, jobID)
// poll.Data.Status: "pending" | "processing" | "completed" | "failed"
// poll.Data.Results available when status == "completed"

// Polling pattern
for {
	poll, err := c.LeadFinder.GetRevealJob(ctx, jobID)
	if err != nil {
		panic(err)
	}
	switch poll.Data.Status {
	case "completed":
		fmt.Println(poll.Data.Results.Revealed)
		return
	case "failed":
		panic(*poll.Data.Error)
	}
	time.Sleep(2 * time.Second)
}

// List recent reveal/enrich jobs
jobs, err := c.LeadFinder.ListRevealJobs(ctx, &sdk.ListRevealJobsRequest{
	Status: "completed",
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
