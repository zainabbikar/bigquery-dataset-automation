# BigQuery Dataset Automation (Golang)

This project demonstrates how to create and delete multiple BigQuery datasets automatically using Go.

## Setup

1. Install Go (>=1.21)
2. Set your service account key:
   ```bash
   export GCP_KEY="path/to/service-account.json"

3. Replace YOUR_PROJECT_ID in main.go
4. Download dependencies:
go mod tidy
5.Run the program:
go run main.go
6. Program will create 10 datasets named demo_dataset_1 to demo_dataset_10 and delete them immediately.