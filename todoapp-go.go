package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TodoappGoStackProps struct {
	awscdk.StackProps
}

func NewTodoappGoStack(scope constructs.Construct, id string, props *TodoappGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// DynamoDB
	table := awsdynamodb.NewTable(stack, jsii.String("TodosTable"), &awsdynamodb.TableProps{
		TableName: jsii.String("todoapp-go-todos-table"),
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// Lambda関数
	function := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("TodoFunction"), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String("todoapp-go-function"),
		Entry:        jsii.String("src"),
		Environment: &map[string]*string{
			"TODOS_TABLE_NAME": table.TableName(),
		},
	})
	// LambdaにDynamoDBのCRUD操作権限を付与
	function.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: jsii.Strings(
			"dynamodb:Scan",
			"dynamodb:PutItem",
			"dynamodb:UpdateItem",
			"dynamodb:DeleteItem"),
		Resources: jsii.Strings(*table.TableArn()),
	}))

	// LogGroup
	awslogs.NewLogGroup(stack, jsii.String("TodoFunctionLogs"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s", *function.FunctionName())),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// API Gateway
	restapi := awsapigateway.NewRestApi(stack, jsii.String("TodoApi"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("todoapp-go-api"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
			StatusCode:   jsii.Number(200),
		},
	})
	todosResource := restapi.
		Root().
		AddResource(jsii.String("todos"), nil)
	todosResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todosResource.AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todoIdResource := todosResource.AddResource(jsii.String("{id}"), nil)
	todoIdResource.AddMethod(jsii.String("PUT"), awsapigateway.NewLambdaIntegration(function, nil), nil)
	todoIdResource.AddMethod(jsii.String("DELETE"), awsapigateway.NewLambdaIntegration(function, nil), nil)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewTodoappGoStack(app, "TodoappGoStack", &TodoappGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
