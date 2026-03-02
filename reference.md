# Reference
## Company Followers
<details><summary><code>client.CompanyFollowers.StartCountEstimate(request) -> *sdk.StartCountEstimateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** — This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Starts an asynchronous count estimation job. Returns the estimated number of followers matching your filters, with optional breakdowns by department and seniority.

**Cost:** 100 credits (fixed). Cached results for the same company + filters within 24 hours are free.

**Results:** Poll the status endpoint to get results. Typically completes within 30–60 seconds.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CountEstimateRequest{
        CompanyURL: "https://linkedin.com/company/stripe",
        Departments: []sdk.CountEstimateRequestDepartmentsItem{
            sdk.CountEstimateRequestDepartmentsItemEngineering,
            sdk.CountEstimateRequestDepartmentsItemSales,
        },
    }
client.CompanyFollowers.StartCountEstimate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyURL:** `string` — The LinkedIn company URL to estimate follower count for
    
</dd>
</dl>

<dl>
<dd>

**countries:** `[]*sdk.CountEstimateRequestCountriesItem` — Filter by country names (57 supported countries)
    
</dd>
</dl>

<dl>
<dd>

**jobTitles:** `[]*sdk.CountEstimateRequestJobTitlesItem` — Filter by job titles (642 supported titles, e.g. Account Executive, Chief Executive Officer, Data Scientist)
    
</dd>
</dl>

<dl>
<dd>

**departments:** `[]*sdk.CountEstimateRequestDepartmentsItem` — Filter by department, returns breakdown for selected departments (26 supported)
    
</dd>
</dl>

<dl>
<dd>

**levels:** `[]*sdk.CountEstimateRequestLevelsItem` — Filter by seniority level, returns breakdown for selected levels (10 supported)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.GetCountEstimateStatus(BatchID) -> *sdk.CountEstimateStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** — This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Check the status of a count estimation job and retrieve results when complete.

**Cost:** Free — polling never costs credits.

**Refunds:** If the job fails within 7 days, the 100 credits are refunded automatically on the next status check.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCountEstimateStatusRequest{
        BatchID: "665b2a3f4c5d6e0013fgh002",
    }
client.CompanyFollowers.GetCountEstimateStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you started the estimation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.StartCompanyFollowerScrape(request) -> *sdk.StartCompanyFollowerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** — This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Starts an asynchronous company follower scrape. Provide a company URL and a max_limit for how many profiles to retrieve.

**Cost:** Credits per profile (see the /limit endpoint for your team's rate), reserved upfront based on max_limit. When you fetch results, we settle the bill — if fewer profiles were found, the excess credits are refunded automatically.

**Filters:** Narrow results by country, job title, department, or seniority level.

**Webhooks:** Optionally pass a `webhookUrl` to automatically receive one webhook per follower profile as it's scraped, plus a final completion callback when the job finishes. No separate flag needed — providing the URL enables per-result webhooks.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CompanyFollowerRequest{
        CompanyURL: "https://linkedin.com/company/stripe",
        MaxLimit: 500,
    }
client.CompanyFollowers.StartCompanyFollowerScrape(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyURL:** `string` — The LinkedIn company URL to scrape followers from
    
</dd>
</dl>

<dl>
<dd>

**maxLimit:** `int` — Maximum number of follower profiles to retrieve (1–50,000)
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Optional custom name for this extraction (visible in history)
    
</dd>
</dl>

<dl>
<dd>

**countries:** `[]*sdk.CompanyFollowerRequestCountriesItem` — Filter followers by country names (57 supported countries)
    
</dd>
</dl>

<dl>
<dd>

**jobTitles:** `[]*sdk.CompanyFollowerRequestJobTitlesItem` — Filter followers by job titles (642 supported titles, e.g. Account Executive, Chief Executive Officer, Data Scientist, Product Manager, Senior Software Engineer)
    
</dd>
</dl>

<dl>
<dd>

**departments:** `[]*sdk.CompanyFollowerRequestDepartmentsItem` — Filter followers by department (26 supported departments)
    
</dd>
</dl>

<dl>
<dd>

**levels:** `[]*sdk.CompanyFollowerRequestLevelsItem` — Filter followers by seniority level (10 supported levels)
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — URL to receive results when scraping completes
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.GetCompanyFollowerProgress(BatchID) -> *sdk.CompanyFollowerProgressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** — This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Poll this endpoint to track the progress of your company follower scrape.

**Cost:** Free — polling never costs credits.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCompanyFollowerProgressRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.CompanyFollowers.GetCompanyFollowerProgress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you started the scrape
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.GetCompanyFollowerResults(BatchID) -> *sdk.CompanyFollowerResultsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Returns paginated follower profiles for a completed batch. Use `page` and `limit` query params to paginate through results.

Credits are charged upfront when you start the scrape. If fewer profiles are found than `max_limit`, unused credits are refunded automatically.

If you provided a `webhookUrl`, results are already delivered via webhooks and credits are auto-settled. No need to call this endpoint.

Only call this after the progress endpoint reports `done: true`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCompanyFollowerResultsRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.CompanyFollowers.GetCompanyFollowerResults(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you started the scrape
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Results per page, max 500 (default: 100)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.ExportCompanyFollowerCsv(BatchID) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Access required** — This endpoint is available to approved accounts only. [Contact us](https://enrich.so) to request access.

Download scrape results as a CSV file. Batch must be complete. No credits consumed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ExportCompanyFollowerCsvRequest{
        BatchID: "batchId",
    }
client.CompanyFollowers.ExportCompanyFollowerCsv(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you started the scrape
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CompanyFollowers.CheckCompanyFollowerLimit() -> *sdk.CheckCompanyFollowerLimitResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check your team's daily company follower scraping limit and current usage. No credits consumed.

**Possible statuses:**
- `approved` — Team has access; `creditsPerProfile`, `dailyLimit`, `dailyUsed`, and `dailyRemaining` are included.
- `pending` — Application is under review.
- `rejected` — Application was rejected.
- `revoked` — Access was revoked.
- `none` — No application on file.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CompanyFollowers.CheckCompanyFollowerLimit(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Email Validation
<details><summary><code>client.EmailValidation.ValidateEmail(request) -> *sdk.EmailValidationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Checks whether an email address is deliverable. The response tells you the
result (`valid`, `invalid`, or `risky`), the confidence level, the mail
provider, and whether the domain is a catch-all.

**Cost:** 1 credit. You are **not** charged if the result is `risky`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.EmailValidationRequest{
        Email: "sarah.chen@stripe.com",
    }
client.EmailValidation.ValidateEmail(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` — The email address you want to validate
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailValidation.BatchValidateEmails(request) -> *sdk.BatchSubmitResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit up to **500 000 emails** in a single request. We automatically
remove duplicates (case-insensitive) before processing, so you only pay
for unique addresses.

**Cost:** 1 credit per unique email, reserved when you submit. Any unused
credits are refunded automatically when you fetch results.

### Webhook callbacks

If you include a `webhookUrl`, your server will receive:

1. **A per-result callback** every time an individual email finishes
   validation. See [emailValidationResult](#tag/Webhooks/operation/webhookEmailValidationResult)
   for the exact payload.

2. **A completion callback** once every email in the batch has been
   processed. See [emailValidationCompletion](#tag/Webhooks/operation/webhookEmailValidationCompletion)
   for the exact payload.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BatchEmailValidationRequest{
        Emails: []string{
            "sarah.chen@stripe.com",
            "james.rodriguez@hubspot.com",
            "priya.patel@notion.so",
            "marco.silva@datadog.com",
        },
        WebhookURL: sdk.String(
            "https://api.yourapp.com/webhooks/enrich",
        ),
    }
client.EmailValidation.BatchValidateEmails(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**emails:** `[]string` — A list of email addresses to validate (up to 500 000)
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` 

A URL on your server where we should send webhook callbacks.
You'll receive one POST per result, plus one final completion POST.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailValidation.GetEmailValidationBatchStatus(BatchID) -> *sdk.BatchStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Poll this endpoint to see how your batch is doing. The `progress` field
gives you a percentage, and `status` will be one of `queued`, `processing`,
`completed`, or `failed`.

**Cost:** Free — polling never costs credits.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetEmailValidationBatchStatusRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.EmailValidation.GetEmailValidationBatchStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailValidation.GetEmailValidationBatchResults(BatchID) -> *sdk.EmailValidationBatchResultsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetch the results once your batch has finished. Results are paginated — use
`page` and `limit` to walk through them.

**Credit settlement:** The first time you call this after the batch reaches
`completed` or `failed`, we calculate the final cost and refund any excess
credits. Subsequent calls return the cached settlement.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetEmailValidationBatchResultsRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.EmailValidation.GetEmailValidationBatchResults(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Results per page (default: 100, max: 1 000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Email Finder
<details><summary><code>client.EmailFinder.FindEmail(request) -> *sdk.EmailFinderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Give us a person's first name, last name, and their company domain, and we'll
find their professional email address.

**Cost:** 10 credits. You are **not** charged if we can't find an email
(`found: false`).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.EmailFinderRequest{
        FirstName: "Emily",
        LastName: "Zhang",
        Domain: "figma.com",
    }
client.EmailFinder.FindEmail(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**firstName:** `string` — The person's first name
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `string` — The person's last name
    
</dd>
</dl>

<dl>
<dd>

**domain:** `string` — The company domain to search (e.g. `figma.com`)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailFinder.BatchFindEmails(request) -> *sdk.BatchSubmitResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit up to **500 000 leads** in one request. Each lead is a combination
of first name, last name, and domain. We deduplicate the list
(case-insensitive) so you are only charged for unique leads.

**Cost:** 10 credits per unique lead, reserved at submission. Unused credits
are refunded when you fetch results.

### Webhook callbacks

If you include a `webhookUrl`, your server will receive:

1. **A per-result callback** each time a single lead is processed. See
   [emailFinderResult](#tag/Webhooks/operation/webhookEmailFinderResult) for the payload.

2. **A completion callback** once the entire batch is done. See
   [emailFinderCompletion](#tag/Webhooks/operation/webhookEmailFinderCompletion) for the payload.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BatchEmailFinderRequest{
        Leads: []*sdk.Lead{
            &sdk.Lead{
                FirstName: "Emily",
                LastName: "Zhang",
                Domain: "figma.com",
            },
            &sdk.Lead{
                FirstName: "David",
                LastName: "Kim",
                Domain: "vercel.com",
            },
            &sdk.Lead{
                FirstName: "Aisha",
                LastName: "Johnson",
                Domain: "cloudflare.com",
            },
        },
        WebhookURL: sdk.String(
            "https://api.yourapp.com/webhooks/enrich",
        ),
    }
client.EmailFinder.BatchFindEmails(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**leads:** `[]*sdk.Lead` — The list of people to search for (up to 500 000)
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` 

A URL on your server where we should send webhook callbacks.
You'll receive one POST per lead, plus one final completion POST.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailFinder.GetEmailFinderBatchStatus(BatchID) -> *sdk.BatchStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Poll this endpoint to track your batch. The `status` moves through
`queued` → `processing` → `completed` (or `failed`).

**Cost:** Free.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetEmailFinderBatchStatusRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.EmailFinder.GetEmailFinderBatchStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailFinder.GetEmailFinderBatchResults(BatchID) -> *sdk.EmailFinderBatchResultsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the results once the batch has finished. Results are paginated.

**Credit settlement:** On the first call after the batch reaches a terminal
status, we settle the final credit charge. Leads where no email was found
are not charged, and the difference is refunded.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetEmailFinderBatchResultsRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.EmailFinder.GetEmailFinderBatchResults(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Results per page (default: 100, max: 1 000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Phone Finder
<details><summary><code>client.PhoneFinder.PhoneLookup() -> *sdk.PhoneLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find phone numbers associated with a person's email address or social
profile URL. You must provide at least one of the two.

**Cost:** 500 credits. You are **not** charged if no phone numbers are found.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PhoneLookupRequest{
        Email: sdk.String(
            "sarah.chen@stripe.com",
        ),
        Linkedin: sdk.String(
            "https://www.example.com/in/sarahchen",
        ),
    }
client.PhoneFinder.PhoneLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `*string` — The person's email address
    
</dd>
</dl>

<dl>
<dd>

**linkedin:** `*string` — The person's profile URL
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneFinder.CreatePhoneBulkJob(request) -> *sdk.PhoneBulkJobSubmitResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit up to **500 000 items** (email addresses and/or profile URLs) for
bulk phone lookup. Emails are deduplicated case-insensitively; profile
URLs are deduplicated as-is.

**Cost:** 10 credits per unique item, reserved upfront. At settlement,
only items where phones were actually found are charged — the rest are
refunded.

### Webhook callbacks

If you include a `webhookUrl`, your server will receive:

1. **A per-result callback** for every item as it is processed. See
   [phoneResult](#tag/Webhooks/operation/webhookPhoneResult) for the payload.

2. **A completion callback** when every item has been processed. See
   [phoneCompletion](#tag/Webhooks/operation/webhookPhoneCompletion) for the payload.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PhoneBatchRequest{
        Emails: []string{
            "sarah.chen@stripe.com",
            "james.rodriguez@hubspot.com",
        },
        Linkedins: []string{
            "https://www.example.com/in/priyapatel",
        },
        WebhookURL: sdk.String(
            "https://api.yourapp.com/webhooks/enrich",
        ),
    }
client.PhoneFinder.CreatePhoneBulkJob(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**emails:** `[]string` — Email addresses to look up
    
</dd>
</dl>

<dl>
<dd>

**linkedins:** `[]string` — Profile URLs to look up
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` 

A URL on your server where we should send webhook callbacks.
You'll receive one POST per item, plus one final completion POST.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneFinder.GetPhoneBulkJobStatus(JobID) -> *sdk.PhoneBulkJobStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Poll this endpoint to track the job. The `progress` field is a percentage
from 0 to 100.

**Cost:** Free.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetPhoneBulkJobStatusRequest{
        JobID: "665c52bf7b8e3f00149f0a10",
    }
client.PhoneFinder.GetPhoneBulkJobStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobID:** `string` — The job identifier (24-character hex string) returned when you submitted the job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneFinder.GetPhoneBulkJobResults(JobID) -> *sdk.PhoneBulkJobResultsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve paginated results. Each item shows the phones found (or an error
message if the lookup failed).

**Credit settlement:** Only items where at least one phone number was found
are charged. Everything else is refunded.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetPhoneBulkJobResultsRequest{
        JobID: "665c52bf7b8e3f00149f0a10",
    }
client.PhoneFinder.GetPhoneBulkJobResults(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobID:** `string` — The job identifier (24-character hex string) returned when you submitted the job
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Results per page (default: 100, max: 1 000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reverse Email Lookup
<details><summary><code>client.ReverseEmailLookup.ReverseLookup(request) -> *sdk.ReverseLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pass an email address and get back the person's professional profile — including
their name, headline, current company, job history, education, and skills.

**Cost:** 10 credits. You are **not** charged if no profile is found.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LookupRequest{
        Email: "emily.zhang@figma.com",
    }
client.ReverseEmailLookup.ReverseLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `string` — The email address to look up
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReverseEmailLookup.BulkReverseLookup(request) -> *sdk.BulkLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit up to **100 000 emails** and get back professional profile data for each
one. Duplicate emails are removed (case-insensitive) before processing.

**Cost:** 10 credits per unique email, reserved upfront. Credits for emails
where no profile was found are refunded.

### Webhook callbacks

If you include a `webhookUrl`, your server will receive:

1. **A per-result callback** for every email as it is processed. See
   [reverseLookupResult](#tag/Webhooks/operation/webhookReverseLookupResult) for the payload.

2. **A completion callback** when the entire batch finishes. See
   [reverseLookupCompletion](#tag/Webhooks/operation/webhookReverseLookupCompletion) for the payload.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkLookupRequest{
        Emails: []string{
            "emily.zhang@figma.com",
            "david.kim@vercel.com",
            "aisha.johnson@cloudflare.com",
        },
        WebhookURL: sdk.String(
            "https://api.yourapp.com/webhooks/enrich",
        ),
    }
client.ReverseEmailLookup.BulkReverseLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**emails:** `[]string` — Email addresses to look up (up to 100 000)
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` 

A URL on your server where we should send webhook callbacks.
You'll receive one POST per email, plus one final completion POST.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReverseEmailLookup.GetBulkLookupStatus(BatchID) -> *sdk.BatchStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Poll this endpoint to track the batch.

**Cost:** Free.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetBulkLookupStatusRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.ReverseEmailLookup.GetBulkLookupStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReverseEmailLookup.GetBulkLookupResults(BatchID) -> *sdk.BulkLookupResultsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve paginated professional profile results. Each result contains the full
profile data, or the result is omitted if no profile was found for that email.

**Credit settlement:** Emails with no matching profile are not charged.
The difference between reserved and actual credits is refunded.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetBulkLookupResultsRequest{
        BatchID: "665a1f4e2c3b7800129dce01",
    }
client.ReverseEmailLookup.GetBulkLookupResults(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchID:** `string` — The batch identifier returned when you submitted the job
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Results per page (default: 100, max: 1 000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Wallets
<details><summary><code>client.Wallets.GetWalletBalance() -> *sdk.WalletBalanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the current credit balance for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Wallets.GetWalletBalance(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Wallets.GetWalletTransactions() -> *sdk.WalletTransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Browse your credit transactions — top-ups, deductions, refunds, and
adjustments. Use the `type` parameter to filter by transaction type.
Results are paginated.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetWalletTransactionsRequest{}
client.Wallets.GetWalletTransactions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page number (default: 1)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Items per page (default: 50, max: 100)
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Filter by type: `deduction`, `topup`, `adjustment`, or `refund`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Teams
<details><summary><code>client.Teams.ListTeamMembers(TeamID) -> *sdk.TeamMemberListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns everyone in your team along with their role (owner, admin, or member).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ListTeamMembersRequest{
        TeamID: "665e0b2f4a6d8c001abc1234",
    }
client.Teams.ListTeamMembers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamID:** `string` — Your team identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Teams.ListTeamInvitations(TeamID) -> *sdk.InvitationListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

See all outstanding invitations for your team. You must be an **admin** or
**owner**.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ListTeamInvitationsRequest{
        TeamID: "665e0b2f4a6d8c001abc1234",
    }
client.Teams.ListTeamInvitations(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamID:** `string` — Your team identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Teams.InviteTeamMember(TeamID, request) -> *sdk.InvitationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sends an email invitation to join your team. You must be an **admin** or
**owner** to invite people. The invitee can be assigned the role of `admin`
or `member`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.InviteBody{
        TeamID: "665e0b2f4a6d8c001abc1234",
        Email: "marco.silva@yourcompany.com",
        Role: sdk.InviteBodyRoleMember,
    }
client.Teams.InviteTeamMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamID:** `string` — Your team identifier
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` — The email address of the person you want to invite
    
</dd>
</dl>

<dl>
<dd>

**role:** `*sdk.InviteBodyRole` — The role to assign them
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Teams.CancelInvitation(TeamID, InvitationID) -> *sdk.SuccessOperationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Withdraws an invitation before the invitee accepts it. You must be an
**admin** or **owner**.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CancelInvitationRequest{
        TeamID: "665e0b2f4a6d8c001abc1234",
        InvitationID: "665e2e7f4a6d8c001abc5001",
    }
client.Teams.CancelInvitation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**teamID:** `string` — Your team identifier
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` — The invitation identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Lead Finder
<details><summary><code>client.LeadFinder.SearchLeads(request) -> *sdk.LeadSearchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search across people, companies, and insights dimensions with comprehensive filters. Returns combined person + org + insights data.

**Free tier:**
- First 75 results free (pages 1–3 at default page size of 25)
- 50 free search queries per month

**Paid tier (page 4+):**
- 1 credit per result returned
- Max 40 pages

Results include preview fields only. Use the **Reveal** or **Enrich** endpoints to get email and phone.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LeadSearchRequest{
        Filters: &sdk.LeadFinderSearchFilters{},
    }
client.LeadFinder.SearchLeads(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**filters:** `*sdk.LeadFinderSearchFilters` — Search filters
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (1–40). Pages 1–3 are free.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Results per page (10–100)
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*string` — Sort field. "none" = primary key order (fastest). Other options include field names like "lastName", "companyName", etc.
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*sdk.LeadSearchRequestSortOrder` — Sort direction
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.CountLeads(request) -> *sdk.LeadCountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the total count of leads matching your filters. This endpoint is free and does not consume credits.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LeadCountRequest{
        Filters: &sdk.LeadFinderSearchFilters{},
    }
client.LeadFinder.CountLeads(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**filters:** `*sdk.LeadFinderSearchFilters` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.ExportLeads(request) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export search results as a CSV file.

**Pricing:**
- Preview fields only: Free
- `includeEmail`: 50 credits per lead
- `includePhone`: 525 credits per lead
- `includeNames`: 1 credit per lead (full last names)
- `includeContactInfo`: Shorthand for email + phone (575 credits/lead)

Previously revealed contacts are returned from cache at no additional cost. Credits are refunded for leads where contact info is unavailable.

**Response:** Returns a CSV file (Content-Type: text/csv) with `X-Credits-Used` and `X-Credits-Refunded` headers.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LeadExportRequest{
        Filters: &sdk.LeadFinderSearchFilters{},
    }
client.LeadFinder.ExportLeads(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**filters:** `*sdk.LeadFinderSearchFilters` 
    
</dd>
</dl>

<dl>
<dd>

**maxResults:** `*int` — Maximum rows to export
    
</dd>
</dl>

<dl>
<dd>

**includeContactInfo:** `*bool` — Shorthand: include both email and phone (costs credits per lead)
    
</dd>
</dl>

<dl>
<dd>

**includeEmail:** `*bool` — Include email addresses (50 credits per lead)
    
</dd>
</dl>

<dl>
<dd>

**includePhone:** `*bool` — Include phone numbers (525 credits per lead)
    
</dd>
</dl>

<dl>
<dd>

**includeNames:** `*bool` — Include full last names (1 credit per lead)
    
</dd>
</dl>

<dl>
<dd>

**leadIDs:** `[]string` — Export specific leads by encrypted ID (selected rows mode)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Export a specific page of results (current page mode)
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Page size when exporting a specific page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.RevealLeads(request) -> *sdk.LeadRevealResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reveal contact info for up to 25 leads. Choose which fields to reveal via the optional `fields` parameter:

- `["email"]` — 50 credits per lead
- `["phone"]` — 525 credits per lead
- `["email", "phone"]` — 575 credits per lead (default if omitted)

Previously revealed contacts are returned from cache at no cost. Returns 402 if insufficient credits.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LeadRevealRequest{
        Leads: []*sdk.LeadRevealRequestLeadsItem{
            &sdk.LeadRevealRequestLeadsItem{
                ID: "enc_abc123def456",
            },
        },
        Fields: []sdk.LeadRevealRequestFieldsItem{
            sdk.LeadRevealRequestFieldsItemEmail,
        },
    }
client.LeadFinder.RevealLeads(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**leads:** `[]*sdk.LeadRevealRequestLeadsItem` — Leads to reveal (max 25)
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*sdk.LeadRevealRequestFieldsItem` — Which contact fields to reveal: 'email' (50 credits), 'phone' (525 credits), or both. Defaults to both if omitted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.EnrichLeads(request) -> *sdk.LeadEnrichResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reveal specific contact fields (email and/or phone) for up to 25 leads.

**Pricing:**
- Email: 50 credits per lead
- Phone: 525 credits per lead

Previously revealed fields are returned from cache at no cost.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.LeadEnrichRequest{
        Leads: []*sdk.LeadEnrichRequestLeadsItem{
            &sdk.LeadEnrichRequestLeadsItem{
                ID: "id",
            },
        },
        Fields: []sdk.LeadEnrichRequestFieldsItem{
            sdk.LeadEnrichRequestFieldsItemEmail,
        },
    }
client.LeadFinder.EnrichLeads(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**leads:** `[]*sdk.LeadEnrichRequestLeadsItem` — Leads to enrich (max 25)
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*sdk.LeadEnrichRequestFieldsItem` — Which contact fields to reveal
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.UnlockNames(request) -> *sdk.UnlockNamesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pay 1 credit per lead to unlock full last names. Search results mask last names by default; this endpoint reveals them.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.UnlockNamesRequest{
        Leads: []*sdk.UnlockNamesRequestLeadsItem{
            &sdk.UnlockNamesRequestLeadsItem{
                ID: "id",
            },
        },
    }
client.LeadFinder.UnlockNames(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**leads:** `[]*sdk.UnlockNamesRequestLeadsItem` — Leads whose last names to unlock (max 100)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.ListSavedSearches() -> *sdk.SavedSearchListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all saved searches for the current user. Free endpoint.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LeadFinder.ListSavedSearches(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.CreateSavedSearch(request) -> *sdk.SavedSearchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Save a search filter combination for reuse. Free endpoint.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SavedSearchCreateRequest{
        Name: "name",
        Filters: &sdk.LeadFinderSearchFilters{},
    }
client.LeadFinder.CreateSavedSearch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Search name
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*sdk.LeadFinderSearchFilters` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.DeleteSavedSearch(ID) -> *sdk.LeadFinderDeletedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a saved search by ID. Only the owner can delete their own saved searches.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DeleteSavedSearchRequest{
        ID: "665e0b2f4a6d8c001abc1234",
    }
client.LeadFinder.DeleteSavedSearch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Saved search ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.GetFilterOptions() -> *sdk.LeadFinderFilterOptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns available values for multi-select filter fields (countries, departments, industries, job functions, etc.). Values are fetched from the database and cached for performance. Free endpoint.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LeadFinder.GetFilterOptions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LeadFinder.SuggestCompanyNames() -> *sdk.LeadFinderSuggestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Prefix-based company name autocomplete. Returns matching company names for search typeahead. Free endpoint.

**Requires approved Lead Finder access.** Returns 403 if your team has not been approved.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SuggestCompanyNamesRequest{
        Q: "Stri",
    }
client.LeadFinder.SuggestCompanyNames(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `string` — Search prefix (min 2 characters)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Max suggestions (1–50, default: 20)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## People Search
<details><summary><code>client.PeopleSearch.WaterfallIcpSearch(request) -> *sdk.WaterfallIcpSearchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for people at a specific company using cascading ICP (Ideal Customer Profile) filter levels.

**How it works:**
1. Provide a company LinkedIn URL to target
2. Optionally define 1-10 cascade levels, each with independent ICP criteria. If omitted, a default 4-level cascade is used: CEO/Founder → C-Suite → VP → Director.
3. Levels are processed in order — results from earlier levels are excluded from later ones
4. Each candidate is scored against the ICP criteria of their matching cascade level
5. Results are returned sorted by score descending

**Scoring:** Each candidate gets a normalized 0-100 score based on title match (30pts), job level (20pts), skills overlap (15pts), location match (15pts), current job status (10pts), profile completeness (5pts), tenure (5pts), and optional seniority priority bonus (0-30pts). The score is normalized against only the dimensions specified in the cascade level.

**Credits:** 1 credit per profile returned. No charge if 0 results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.WaterfallIcpSearchRequest{
        CompanyLinkedinURL: "https://www.linkedin.com/company/google/",
    }
client.PeopleSearch.WaterfallIcpSearch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyLinkedinURL:** `string` — Company LinkedIn URL to search employees for. Must be a valid LinkedIn company URL.
    
</dd>
</dl>

<dl>
<dd>

**cascade:** `[]*sdk.WaterfallIcpSearchRequestCascadeItem` — Cascading ICP filter levels, processed in order. Results from earlier levels are excluded from later ones (deduplication by RBID_PER). If omitted, uses a default 4-level cascade: CEO/Founder → C-Suite → VP → Director.
    
</dd>
</dl>

<dl>
<dd>

**maxResults:** `*int` — Maximum total results to return across all cascade levels
    
</dd>
</dl>

<dl>
<dd>

**minScore:** `*float64` — Minimum ICP score threshold (0-100). Candidates scoring below this are filtered out.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PeopleSearch.EmployeeFinder(request) -> *sdk.EmployeeFinderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find employees at a specific company by LinkedIn URL. Optionally filter by job level, job function, country, continent, or sales region.

**Credits:** 1 credit per result returned. No charge if 0 results.

**Pagination:** Use `page` and `max_results` to paginate through results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.EmployeeFinderRequest{
        CompanyLinkedinURL: "https://www.linkedin.com/company/google/",
    }
client.PeopleSearch.EmployeeFinder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyLinkedinURL:** `string` — LinkedIn company profile URL (required). Must be a valid LinkedIn company URL.
    
</dd>
</dl>

<dl>
<dd>

**country:** `[]string` — Country names to filter (e.g. "India", "United States", "Germany")
    
</dd>
</dl>

<dl>
<dd>

**continent:** `[]string` — Continent filter
    
</dd>
</dl>

<dl>
<dd>

**salesRegion:** `[]string` — Sales region filter
    
</dd>
</dl>

<dl>
<dd>

**jobLevel:** `[]string` — Job level/seniority filter
    
</dd>
</dl>

<dl>
<dd>

**jobFunction:** `[]string` — Job function/department filter
    
</dd>
</dl>

<dl>
<dd>

**maxResults:** `*int` — Results per page (1-100, default 10)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number (starts at 1)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

