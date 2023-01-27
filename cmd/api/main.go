package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/IgorRamos/fm-transaction/configs"
	"github.com/IgorRamos/fm-transaction/internal/handlers"
	"github.com/IgorRamos/fm-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type API struct {
	handler handlers.TransactionHandler
}

func main() {
	appConfig := configs.GetAppConfigs()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(appConfig.DynamoRegion),
	)
	if err != nil {
		log.Fatal("Failed to create aws client, error: ", err.Error())
	}

	db := dynamodb.NewFromConfig(cfg)
	transactionRepository := repositories.NewTransactionRepository(db, appConfig.DynamoTransactionTable)
	transactionHandler := handlers.NewTransactionHandler(transactionRepository)
	api := API{transactionHandler}

	http.HandleFunc("/transactions", api.handleTransactions)
	http.HandleFunc("/report", api.handleReport)

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h API) handleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		req := createRequest(r)
		resp, err := h.handler.GetTransactions(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, resp.Body)
	}

	if r.Method == "POST" {
		req := createRequest(r)
		resp, err := h.handler.CreateTransaction(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, resp.Body)
	}
}

func (h API) handleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		req := createRequest(r)
		resp, err := h.handler.ReportTransactions(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, resp.Body)
	}
}

func createRequest(r *http.Request) events.APIGatewayProxyRequest {
	params := r.URL.Query()
	queryParams := make(map[string]string)
	for k, v := range params {
		queryParams[k] = v[0]
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	req := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         map[string]string{},
		MultiValueHeaders:               map[string][]string{},
		QueryStringParameters:           queryParams,
		MultiValueQueryStringParameters: params,
		PathParameters:                  map[string]string{},
		StageVariables:                  map[string]string{},
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            string(bodyBytes),
		IsBase64Encoded:                 false,
	}
	return req
}
