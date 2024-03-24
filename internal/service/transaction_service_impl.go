package service

import (
	"mime/multipart"
	"stori-lambda/internal/config"
	"stori-lambda/internal/domain"
	"stori-lambda/internal/dto"
	"stori-lambda/internal/mapper"
	"stori-lambda/internal/repository"
)

type TransactionServiceImpl struct {
	TransactionDetailsRepository repository.TransactionDetailsRepository
	smtpClient                   *config.SmtpClient
	aws                          *config.Aws
}

func NewTransactionServiceImpl(TransactionDetailsRepository repository.TransactionDetailsRepository, smtpClient *config.SmtpClient, aws *config.Aws) TransactionService {
	return &TransactionServiceImpl{
		TransactionDetailsRepository: TransactionDetailsRepository,
		smtpClient:                   smtpClient,
		aws:                          aws,
	}
}

func (t TransactionServiceImpl) LoadFileTransaction(loadFile dto.LoadFileRequest, file multipart.File, fileName string) {
	t.aws.PutObjet(loadFile.Path, file)
}

func (t TransactionServiceImpl) SendAccountReport(accountReportRequest dto.AccountReportRequest) {
	var transactionFileRecord = []dto.TransactionFile{}
	t.aws.GetFile(accountReportRequest.Path, &transactionFileRecord)
	go t.saveInfoFile(transactionFileRecord)
	transactionReport := domain.NewTransactionReport(transactionFileRecord)
	bodyMail := t.smtpClient.ParseTemplate("transaction_report.html", transactionReport)
	t.smtpClient.SendMail(accountReportRequest.Mail, bodyMail)
}

func (t TransactionServiceImpl) saveInfoFile(transactionFileRecord []dto.TransactionFile) {
	for _, transactionRecord := range transactionFileRecord {
		transactionDetail := mapper.From(transactionRecord)
		t.TransactionDetailsRepository.Save(transactionDetail)
	}
}
