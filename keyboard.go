package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)


type Keyboard struct {
	Type string `json:"type"`
}
// /keyboard	GET	키보드 영역에 표현될 버튼에 대한 정보.
func keyboardHandler(request events.APIGatewayProxyRequest) (Keyboard, error){
	return Keyboard{
		Type: "text",
	}, nil
}

func main() {
	lambda.Start(keyboardHandler)
}