package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

func (h TransactionHandler) CreateTransaction(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transaction models.Transaction
	err := json.Unmarshal([]byte(req.Body), &transaction)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	err = h.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}
