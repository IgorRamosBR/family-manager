package utils

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func CreateErrorResponse(statusCode int, err error) events.APIGatewayProxyResponse {
	errorResponse := ErrorResponse{
		Code:    strconv.Itoa(statusCode),
		Message: err.Error(),
	}

	errorBody, _ := json.Marshal(errorResponse)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(errorBody),
		Headers:    map[string]string{"Content-type": "application/json"},
	}
}
