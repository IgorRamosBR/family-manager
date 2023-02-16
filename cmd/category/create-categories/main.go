package main

import (
	"context"
	"log"

	"github.com/IgorRamos/fm-transaction/configs"
	handlers "github.com/IgorRamos/fm-transaction/internal/handlers/category"
	"github.com/IgorRamos/fm-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	appConfig := configs.GetAppConfigs()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(appConfig.DynamoRegion),
	)
	if err != nil {
		log.Fatal("Failed to create aws client, error: ", err.Error())
	}

	db := dynamodb.NewFromConfig(cfg)
	categoryRepository := repositories.NewCategoryRepository(db, appConfig.DynamoCategoryTable)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	lambda.Start(categoryHandler.CreateCategory)
}
