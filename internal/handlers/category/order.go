package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

func (h CategoryHandler) UpdateCategoryListOrder(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var categories []models.Category
	err := json.Unmarshal([]byte(req.Body), &categories)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	err = h.categoryRepository.UpdateCategoryListOrder(categories)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
