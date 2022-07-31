package main

import (
	"encoding/json"
	"todoapp-go/src/controllers"
	"todoapp-go/src/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ConvertToString(value interface{}) string {
	bytes, _ := json.Marshal(value)
	return string(bytes)
}

func GetHandler() string {
	response := controllers.GetTodo()
	return ConvertToString(response)
}

func PostHandler(body string) string {
	request := models.NewPostTodoRequestFromString(body)
	response := controllers.PostTodo(request.Message)
	return ConvertToString(response)
}

func PutHandler(id string, body string) string {
	request := models.NewPutTodoRequestFromString(body)
	response := controllers.PutTodo(id, request.Message)
	return ConvertToString(response)
}

func DeleteHandler(id string) string {
	response := controllers.DeleteTodo(id)
	return ConvertToString(response)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	method := request.HTTPMethod
	todoId := request.PathParameters["id"]
	body := request.Body
	var response string

	if method == "GET" {
		response = GetHandler()
	} else if method == "POST" {
		response = PostHandler(body)
	} else if method == "PUT" {
		response = PutHandler(todoId, body)
	} else if method == "DELETE" {
		response = DeleteHandler(todoId)
	}

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
