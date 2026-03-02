package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdk "github.com/maximiseai/enrich-go-sdk"
	"github.com/maximiseai/enrich-go-sdk/client"
	"github.com/maximiseai/enrich-go-sdk/option"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	ctx := context.Background()

	apiKey := getEnv("ENRICH_API_KEY", "")
	baseURL := getEnv("ENRICH_BASE_URL", "https://api.enrich.so/v3")

	if apiKey == "" {
		log.Fatal("ENRICH_API_KEY environment variable is required")
	}

	c := client.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL),
	)

	fmt.Println("=== Enrich SDK – local server tests ===")
	fmt.Println()

	// 1. Wallet balance (free)
	fmt.Println("--- Wallet balance ---")
	balance, err := c.Wallets.GetWalletBalance(ctx)
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", balance)
	}

	// 2. Wallet transactions (free)
	fmt.Println("--- Wallet transactions ---")
	page := 1
	limit := 5
	transactions, err := c.Wallets.GetWalletTransactions(ctx, &sdk.GetWalletTransactionsRequest{
		Page:  &page,
		Limit: &limit,
	})
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", transactions)
	}

	// 3. Lead finder filter options (free)
	fmt.Println("--- Lead finder filter options ---")
	filterOpts, err := c.LeadFinder.GetFilterOptions(ctx)
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %d filter keys returned\n", len(filterOpts.Data))
		for k, v := range filterOpts.Data {
			fmt.Printf("  %s (%s): %d values\n", k, v.Label, len(v.Values))
		}
		fmt.Println()
	}

	// 4. Company suggest (free)
	fmt.Println("--- Company suggest 'Fig' ---")
	suggestions, err := c.LeadFinder.SuggestCompanyNames(ctx, &sdk.SuggestCompanyNamesRequest{
		Q: "Fig",
	})
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", suggestions)
	}

	// 4. Company follower limit (free)
	fmt.Println("--- Company follower limit ---")
	flimit, err := c.CompanyFollowers.CheckCompanyFollowerLimit(ctx)
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", flimit)
	}

	// 5. Email validation (1 credit)
	fmt.Println("--- Email validation ---")
	validation, err := c.EmailValidation.ValidateEmail(ctx, &sdk.EmailValidationRequest{
		Email: "emily@figma.com",
	})
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", validation)
	}

	// 6. Email finder (10 credits)
	fmt.Println("--- Email finder ---")
	found, err := c.EmailFinder.FindEmail(ctx, &sdk.EmailFinderRequest{
		FirstName: "Emily",
		LastName:  "Zhang",
		Domain:    "figma.com",
	})
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", found)
	}

	// 7. Reverse email lookup (10 credits)
	fmt.Println("--- Reverse email lookup ---")
	profile, err := c.ReverseEmailLookup.ReverseLookup(ctx, &sdk.LookupRequest{
		Email: "emily@figma.com",
	})
	if err != nil {
		log.Printf("FAIL: %v\n\n", err)
	} else {
		fmt.Printf("OK: %+v\n\n", profile)
	}

	fmt.Println("=== Done ===")
}
