package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"stori-lambda/internal/dto"
	"stori-lambda/internal/helper"
	"stori-lambda/internal/service"
)

type TransactionController struct {
	service service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{service: transactionService}
}

func (controller *TransactionController) LoadFile(ctx *gin.Context) {
	var request dto.LoadFileRequest
	data := ctx.PostForm("data")
	err := json.Unmarshal([]byte(data), &request)
	_, file, err := ctx.Request.FormFile("file")
	helper.PrintlnError(err, "fallo recibir archivo")
	helper.ErrorPanic(err)
	f, err := file.Open()
	helper.ErrorPanic(err)
	defer f.Close()
	controller.service.LoadFileTransaction(request, f, file.Filename)
	ctx.JSON(http.StatusCreated, "ok")
}

func (contoller TransactionController) SendMail(ctx *gin.Context) {
	request := dto.AccountReportRequest{}
	err := ctx.ShouldBindJSON(&request)
	helper.ErrorPanic(err)
	contoller.service.SendAccountReport(request)
}

func (contoller TransactionController) helpCheck(ctx *gin.Context) {
	ctx.String(200, "Hola mundo")
}
