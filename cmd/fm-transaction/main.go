package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/IgorRamos/fm-create-transaction/configs"
	"github.com/IgorRamos/fm-create-transaction/internal/handlers"
	"github.com/IgorRamos/fm-create-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var transactionHandler handlers.TransactionHandler
var appConfig configs.AppConfig

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Path == "/transactions" {
		if req.HTTPMethod == "POST" {
			return transactionHandler.CreateTransaction(req)
		}
		if req.HTTPMethod == "GET" {
			return transactionHandler.GetTransactions(req)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func main() {
	initComponents()

	if appConfig.Environment == "local" {
		runLocal()
	} else {
		lambda.Start(router)
	}
}

func runLocal() {
	http.HandleFunc("/", mapRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func mapRequest(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	params := make(map[string]string)
	for k, v := range values {
		params[k] = strings.Join(v, ",")
	}

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)

	reqUri := strings.Split(r.RequestURI, "?")
	req := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            reqUri[0],
		HTTPMethod:                      r.Method,
		Headers:                         map[string]string{},
		QueryStringParameters:           params,
		MultiValueQueryStringParameters: map[string][]string{},
		PathParameters:                  map[string]string{},
		StageVariables:                  map[string]string{},
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            string(body),
		IsBase64Encoded:                 false,
	}

	resp, err := router(req)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(resp.Body))
}

func initComponents() {
	appConfig = configs.GetAppConfigs()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(appConfig.DynamoRegion),
	)
	if err != nil {
		log.Fatal("Failed to create aws client, error: ", err.Error())
	}

	db := dynamodb.NewFromConfig(cfg)
	transactionRepository := repositories.NewTransactionRepository(db, appConfig.DynamoTransactionTable)
	transactionHandler = handlers.NewTransactionHandler(transactionRepository)
}
