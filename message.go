package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	User_key        string `json:"user_key"`
	Type        string `json:"type"`
	Content        string `json:"content"`
}
type Message struct {
	Message     map[string]string  `json:"message"`
	Keyboard    map[string]string `json:"keyboard"`
}
// /message	POST	자동응답 명령어에 대한 응답 메시지의 내용
func MessageHandler(request Request) (Message, error) {
	return Message{
		Message: map[string]string{"text": "User_key="+request.User_key+" Type="+request.Type+" Content="+request.Content},
		Keyboard: map[string]string{"type": "text"},
	}, nil
}

func main() {
	lambda.Start(MessageHandler)
}