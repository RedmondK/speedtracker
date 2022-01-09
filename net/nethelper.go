package speedtrackertypes

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func CreateErrorResponse(errorString string) (errorResponse events.APIGatewayV2HTTPResponse) {
	return events.APIGatewayV2HTTPResponse{
		Body:       CreateErrorBody(errorString),
		StatusCode: 500,
	}
}

func CreateErrorBody(errorString string) (ErrorBody string) {
	return fmt.Sprintf(`{"error": "%s"}`, errorString)
}
