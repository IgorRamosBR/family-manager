package handlers

import (
	"net/http"

	"github.com/IgorRamos/fm-create-transaction/internal/repositories"
	"github.com/aws/aws-lambda-go/events"
)

func (h TransactionHandler) GetTransactions(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	period := req.QueryStringParameters["period"]
	if period == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Query Parameter [period] required",
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	offset := req.QueryStringParameters["offset"]
	limit, err := toInt32(req.QueryStringParameters["limit"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Limit must be a number",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	queryParams := repositories.QueryParameters{
		Period: period,
		Offset: offset,
		Limit:  limit,
	}

	page, err := h.transactionRepository.GetTransactions(queryParams)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := toJSON(page)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       responseBody,
		Headers:    map[string]string{"Content-type": "application/json"},
	}, nil
}
