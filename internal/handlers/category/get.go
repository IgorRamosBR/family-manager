package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

func (h CategoryHandler) GetCategories(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	categories, err := h.categoryRepository.GetAllCategories()
	if err != nil {
		log.Errorf("error to get categories, error: [%s]", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := toJSON(categories)
	if err != nil {
		log.Errorf("error to parse JSON, error: [%s]", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       responseBody,
		Headers: map[string]string{
			"Content-type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
