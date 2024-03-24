package service

import (
	"mime/multipart"
	"stori-lambda/internal/dto"
)

type TransactionService interface {
	LoadFileTransaction(loadFile dto.LoadFileRequest, file multipart.File, fileName string)
	SendAccountReport(accountReportRequest dto.AccountReportRequest)
}
