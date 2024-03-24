// cmd/lambda/loadFile/main.go

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"stori-lambda/internal/config"
	"stori-lambda/internal/controller"
	"stori-lambda/internal/repository"
	"stori-lambda/internal/service"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	db := config.DatabaseConnection()
	config.Migration(db)
	aws := config.NewAwsClient()
	transactionDetailsRepository := repository.NewTransactionDetailsRepositoryImpl(db)
	smtp := config.NewSmtpClient()
	transactionService := service.NewTransactionServiceImpl(transactionDetailsRepository, smtp, aws)
	transactionController := controller.NewTransactionController(transactionService)
	r.POST("/loadFile", transactionController.LoadFile)
	r.POST("/sendMail", transactionController.SendMail)
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
