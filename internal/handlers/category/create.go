package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

func (h CategoryHandler) CreateCategory(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var category models.Category
	err := json.Unmarshal([]byte(req.Body), &category)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	err = h.categoryRepository.CreateCategory(category)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
