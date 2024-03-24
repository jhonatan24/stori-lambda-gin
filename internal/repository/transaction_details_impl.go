package repository

import (
	"gorm.io/gorm"
	"stori-lambda/internal/entity"
)

type TransactionDetailsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionDetailsRepositoryImpl(Db *gorm.DB) TransactionDetailsRepository {
	return &TransactionDetailsRepositoryImpl{Db: Db}
}

func (t TransactionDetailsRepositoryImpl) Save(transactionDetails entity.TransactionDetailsEntity) {
	t.Db.Create(&transactionDetails)
}
