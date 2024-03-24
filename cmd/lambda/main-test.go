package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"stori-lambda/internal/config"
	"stori-lambda/internal/controller"
	"stori-lambda/internal/repository"
	"stori-lambda/internal/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	db := config.DatabaseConnection()
	config.Migration(db)
	aws := config.NewAwsClient()
	transactionDetailsRepository := repository.NewTransactionDetailsRepositoryImpl(db)
	smtp := config.NewSmtpClient()
	transactionService := service.NewTransactionServiceImpl(transactionDetailsRepository, smtp, aws)
	transactionController := controller.NewTransactionController(transactionService)
	router := gin.Default()
	router.POST("/loadFile", transactionController.LoadFile)
	router.POST("/sendMail", transactionController.SendMail)
	router.Run(":8080")
}
