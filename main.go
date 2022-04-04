package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "kouzoh-p-bharath"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func main() {
	g := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	// to check if the client will use the right credentials
	fmt.Printf("Path: %s ", g)

	ctx := context.Background()

	client := createClient(ctx)

	operate(ctx, client)

	defer client.Close()
}

func operate(ctx context.Context, client *firestore.Client) error {
	return nil
}
