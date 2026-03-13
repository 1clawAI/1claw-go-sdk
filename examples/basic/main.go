// Example: basic SDK usage with API key and token auth.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/1clawAI/1claw-go-sdk"
)

func main() {
	// Option 1: API key (auto-exchanges for JWT on first call)
	client, err := oneclaw.New(
		oneclaw.WithBaseURL("https://api.1claw.xyz"),
		oneclaw.WithAPIKey(os.Getenv("1CLAW_API_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// List vaults
	vaults, err := client.Vaults.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Vaults: %d\n", len(vaults.Vaults))
	for _, v := range vaults.Vaults {
		fmt.Printf("  - %s (%s)\n", v.Name, v.ID)
	}

	// Option 2: Pre-obtained JWT
	tokenClient, _ := oneclaw.New(
		oneclaw.WithBaseURL("https://api.1claw.xyz"),
		oneclaw.WithToken(os.Getenv("1CLAW_TOKEN")),
	)
	_ = tokenClient
}
