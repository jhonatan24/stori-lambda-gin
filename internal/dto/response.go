package dto

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func APIResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	bytes, _ := json.Marshal(&body)

	return events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: statusCode,
	}, nil
}

func APIErrResponse(statusCode int, err error) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: statusCode,
	}, nil
}

func APIServerError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Internal server error",
		StatusCode: 500,
	}, err
}
