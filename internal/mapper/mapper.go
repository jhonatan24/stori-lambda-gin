package mapper

import (
	"stori-lambda/internal/dto"
	"stori-lambda/internal/entity"
	"strconv"
)

func From(transaction dto.TransactionFile) entity.TransactionDetailsEntity {
	amount, _ := strconv.ParseFloat(transaction.Amount, 64)

	return entity.TransactionDetailsEntity{
		Date:   transaction.GetDateFormat(),
		Amount: amount,
	}
}
