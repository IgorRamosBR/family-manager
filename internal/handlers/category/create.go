package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorRamos/fm-transaction/internal/handlers/utils"
	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

func (h CategoryHandler) CreateCategory(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var categoryRequest models.CategoryRequest
	err := json.Unmarshal([]byte(req.Body), &categoryRequest)
	if err != nil {
		return utils.CreateErrorResponse(http.StatusBadRequest, err), nil

	}

	err = categoryRequest.Validate()
	if err != nil {
		return utils.CreateErrorResponse(http.StatusBadRequest, err), nil
	}

	category := categoryRequest.ToDynamoModel()

	err = h.categoryRepository.CreateCategory(category)
	if err != nil {
		return utils.CreateErrorResponse(http.StatusInternalServerError, err), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
