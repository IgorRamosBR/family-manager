package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/IgorRamos/fm-transaction/configs"
	categoryHandler "github.com/IgorRamos/fm-transaction/internal/handlers/category"
	transactionHandler "github.com/IgorRamos/fm-transaction/internal/handlers/transaction"
	"github.com/IgorRamos/fm-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type API struct {
	transactionHandler transactionHandler.TransactionHandler
	categoryHandler    categoryHandler.CategoryHandler
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
	categoryRepository := repositories.NewCategoryRepository(db, appConfig.DynamoCategoryTable)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryRepository)
	transactionRepository := repositories.NewTransactionRepository(db, appConfig.DynamoTransactionTable)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionRepository, categoryRepository)

	api := API{transactionHandler: transactionHandler, categoryHandler: categoryHandler}

	http.HandleFunc("/transactions", api.handleTransactions)
	http.HandleFunc("/report", api.handleReport)
	http.HandleFunc("/categories", api.handleCategories)
	http.HandleFunc("/categories/order", api.handleCategoriesOrder)

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h API) handleTransactions(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		setupCORS(&w, r)
		return
	}
	if r.Method == "GET" {
		req := createRequest(r)
		resp, err := h.transactionHandler.GetTransactions(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		for k, v := range resp.Headers {
			w.Header().Add(k, v)
		}
		io.WriteString(w, resp.Body)
	}

	if r.Method == "POST" {
		req := createRequest(r)
		resp, err := h.transactionHandler.CreateTransaction(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		for k, v := range resp.Headers {
			w.Header().Add(k, v)
		}
		io.WriteString(w, resp.Body)
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (h API) handleReport(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "GET" {
		req := createRequest(r)
		resp, err := h.transactionHandler.ReportTransactions(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, resp.Body)
	}
}

func (h API) handleCategories(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		setupCORS(&w, r)
		return
	}
	if r.Method == "GET" {
		req := createRequest(r)
		resp, err := h.categoryHandler.GetCategories(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		for k, v := range resp.Headers {
			w.Header().Add(k, v)
		}
		io.WriteString(w, resp.Body)
	}
	if r.Method == "POST" {
		req := createRequest(r)
		resp, err := h.categoryHandler.CreateCategory(req)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		io.WriteString(w, resp.Body)
	}
}

func (h API) handleCategoriesOrder(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		req := createRequest(r)
		resp, err := h.categoryHandler.UpdateCategoryListOrder(req)
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

	headers := map[string]string{}
	for k, v := range r.Header {
		headers[k] = v[0]
	}

	req := events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         headers,
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
