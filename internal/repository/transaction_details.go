package repository

import "stori-lambda/internal/entity"

type TransactionDetailsRepository interface {
	Save(transactionDetails entity.TransactionDetailsEntity)
}
