package handlers

import (
	"context"
	"log"
	"testing"

	"github.com/IgorRamos/fm-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func TestUpdateCategoryListOrder(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKIARKEFHTPIZXGLDHWD", "2DWuJ3ydZ8MZvVLPud+IqFcozuWUDbZf/3NkcEri", "")),
	)
	if err != nil {
		log.Fatal("Failed to create aws client, error: ", err.Error())
	}

	db := dynamodb.NewFromConfig(cfg)
	categoryRepository := repositories.NewCategoryRepository(db, "fm-category-table-dev")
	categoryHandler := NewCategoryHandler(categoryRepository)

	req := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         map[string]string{},
		MultiValueHeaders:               map[string][]string{},
		QueryStringParameters:           map[string]string{},
		MultiValueQueryStringParameters: map[string][]string{},
		PathParameters:                  map[string]string{},
		StageVariables:                  map[string]string{},
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            `[{"name":"Habitação - Prestação","type":"EXPENSE","priority":1},{"name":"Habitação - Condomínio","type":"EXPENSE","priority":2},{"name":"Habitação - Internet","type":"EXPENSE","priority":3},{"name":"Habitação - Luz","type":"EXPENSE","priority":4}]`,
		IsBase64Encoded:                 false,
	}

	categoryHandler.UpdateCategoryListOrder(req)
}
