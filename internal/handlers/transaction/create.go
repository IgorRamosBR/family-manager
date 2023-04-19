package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorRamos/fm-transaction/internal/handlers/utils"
	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

func (h TransactionHandler) CreateTransaction(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transactionRequest models.TransactionRequest
	err := json.Unmarshal([]byte(req.Body), &transactionRequest)
	if err != nil {
		return utils.CreateErrorResponse(http.StatusBadRequest, err), nil
	}

	err = transactionRequest.Validate()
	if err != nil {
		return utils.CreateErrorResponse(http.StatusBadRequest, err), nil
	}

	category, err := h.categoryRepository.GetCategory(transactionRequest.Category.Name, transactionRequest.Category.Priority)
	if err != nil {
		return utils.CreateErrorResponse(http.StatusBadRequest, err), nil
	}

	transaction := transactionRequest.ToDynamoModel(category)

	err = h.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return utils.CreateErrorResponse(http.StatusInternalServerError, err), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}
