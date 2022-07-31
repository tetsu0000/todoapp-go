package controllers

import (
	"context"
	"os"
	"todoapp-go/src/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

var ctx = context.Background()
var region = os.Getenv("AWS_REGION")
var cfg, _ = config.LoadDefaultConfig(ctx, config.WithRegion(region))
var client = dynamodb.NewFromConfig(cfg)
var tableName = os.Getenv("TODOS_TABLE_NAME")

func GetTodo() *models.GetTodosResponseBody {
	scanResult, _ := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: &tableName,
	})
	todos := new(models.Todos)
	attributevalue.UnmarshalListOfMaps(scanResult.Items, todos)
	return models.NewGetTodoResponseBody("Get todos succeed.", todos)
}

func PostTodo(message string) *models.PostTodoResponseBody {
	uuidIns, _ := uuid.NewRandom()
	id := uuidIns.String()
	client.PutItem(ctx, &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: id},
			"message": &types.AttributeValueMemberS{Value: message},
		},
		TableName: &tableName,
	})
	todo := models.NewTodo(id, message)
	return models.NewPostTodoResponseBody("Post todo succeed.", todo)
}

func PutTodo(id string, message string) *models.PutTodoResponseBody {
	client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
		UpdateExpression: aws.String("SET message = :m"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":m": &types.AttributeValueMemberS{
				Value: message,
			},
		},
	})
	todo := models.NewTodo(id, message)
	return models.NewPutTodoResponseBody("Put todo succeed.", todo)
}

func DeleteTodo(id string) *models.DeleteTodoResponseBody {
	client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	})
	return models.NewDeleteTodoResponseBody("Delete todo succeed.")
}
