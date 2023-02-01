package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Environment            string
	DynamoTransactionTable string
	DynamoCategoryTable    string
	DynamoEndpoint         string
	DynamoRegion           string
}

func GetAppConfigs() AppConfig {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Fatal("Failed to read ENVIRONMENT")
	}

	dynamoTransactionTable := os.Getenv("DYNAMODB_TRANSACTION_TABLE")
	if dynamoTransactionTable == "" {
		log.Fatal("Failed to read DYNAMODB_TRANSACTION_TABLE")
	}

	dynamoCategoryTable := os.Getenv("DYNAMODB_CATEGORY_TABLE")
	if dynamoCategoryTable == "" {
		log.Fatal("Failed to read DYNAMODB_CATEGORY_TABLE")
	}

	dynamoEndpoint := os.Getenv("DYNAMODB_TRANSACTION_ENDPOINT")
	if dynamoEndpoint == "" {
		log.Fatal("Failed to read DYNAMODB_TRANSACTION_ENDPOINT")
	}

	dynamoRegion := os.Getenv("DYNAMODB_TRANSACTION_REGION")
	if dynamoRegion == "" {
		log.Fatal("Failed to read DYNAMODB_TRANSACTION_REGION")
	}

	return AppConfig{
		Environment:            environment,
		DynamoTransactionTable: dynamoTransactionTable,
		DynamoCategoryTable:    dynamoCategoryTable,
		DynamoEndpoint:         dynamoEndpoint,
		DynamoRegion:           dynamoRegion,
	}
}
