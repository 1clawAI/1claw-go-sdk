// Example: CMEK client-side encryption.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/1clawAI/1claw-go-sdk"
	"github.com/1clawAI/1claw-go-sdk/cmek"
)

func main() {
	client, err := oneclaw.New(
		oneclaw.WithBaseURL("https://api.1claw.xyz"),
		oneclaw.WithAPIKey(os.Getenv("1CLAW_API_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Generate CMEK key
	key, err := cmek.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fingerprint := cmek.Fingerprint(key)
	fmt.Printf("Fingerprint: %s\n", fingerprint)

	// Encrypt before storing
	plaintext := "my-api-key-secret"
	encrypted, err := cmek.Encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	vaultID := os.Getenv("1CLAW_VAULT_ID")
	if vaultID == "" {
		log.Println("Set 1CLAW_VAULT_ID to run store. Skipping.")
		return
	}

	_, err = client.Secrets.PutSecret(ctx, vaultID, "cmek/encrypted", encrypted, "generic")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stored encrypted secret")

	// Retrieve and decrypt
	secret, err := client.Secrets.GetSecret(ctx, vaultID, "cmek/encrypted")
	if err != nil {
		log.Fatal(err)
	}
	decrypted, err := cmek.Decrypt(key, secret.Value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
