package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Set your service account JSON path in env variable
	keyPath := os.Getenv("GCP_KEY")
	if keyPath == "" {
		log.Fatal("GCP_KEY environment variable not set")
	}

	projectID := "YOUR_PROJECT_ID"

	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Create and delete 10 datasets
	for i := 1; i <= 10; i++ {
		datasetID := fmt.Sprintf("demo_dataset_%d", i)

		// Create dataset
		dataset := client.Dataset(datasetID)
		if err := dataset.Create(ctx, &bigquery.DatasetMetadata{
			Location: "US",
		}); err != nil {
			log.Fatalf("Failed to create dataset %s: %v", datasetID, err)
		}
		fmt.Printf("Created dataset: %s\n", datasetID)

		// Optional: wait 1 second between operations
		time.Sleep(1 * time.Second)

		// Delete dataset
		if err := dataset.Delete(ctx); err != nil {
			log.Fatalf("Failed to delete dataset %s: %v", datasetID, err)
		}
		fmt.Printf("Deleted dataset: %s\n", datasetID)
	}
}
