package domain

import (
	"stori-lambda/internal/dto"
)

type TransactionReport struct {
	TotalAmount                float64
	CountTransactionByMouth    map[string]int
	AverageDebitAmountByMouth  map[string]float64
	AverageCreditAmountByMouth map[string]float64
}

func NewTransactionReport(transactionFiles []dto.TransactionFile) TransactionReport {
	var totalAmount float64 = 0
	var countTransactionByMouth = make(map[string]int)
	var listDebitAmountByMouth = make(map[string][]float64)
	var listCreditAmountByMouth = make(map[string][]float64)

	for _, transactionFileRow := range transactionFiles {
		totalAmount += transactionFileRow.GetAmountFloat()
		countByMouth := countTransactionByMouth[transactionFileRow.GetMoth().String()]
		countByMouth++
		countTransactionByMouth[transactionFileRow.GetMoth().String()] = countByMouth

		appendMap(listDebitAmountByMouth, transactionFileRow, transactionFileRow.IsDebit())
		appendMap(listCreditAmountByMouth, transactionFileRow, !transactionFileRow.IsDebit())
	}
	return TransactionReport{
		TotalAmount:                totalAmount,
		CountTransactionByMouth:    countTransactionByMouth,
		AverageCreditAmountByMouth: averageMapByKey(listCreditAmountByMouth),
		AverageDebitAmountByMouth:  averageMapByKey(listDebitAmountByMouth),
	}
}

func appendMap(elem map[string][]float64, transactionFileRow dto.TransactionFile, condition bool) {
	if !condition {
		return
	}
	debitsAmounts, _ := elem[transactionFileRow.GetMoth().String()]
	debitsAmounts = append(debitsAmounts, transactionFileRow.GetAmountFloat())
	elem[transactionFileRow.GetMoth().String()] = debitsAmounts
}

func averageMapByKey(elem map[string][]float64) map[string]float64 {
	averageMapBykey := make(map[string]float64)
	for key, values := range elem {
		sumValue := sumArray(values)
		average := sumValue / float64(len(values))
		averageMapBykey[key] = average
	}
	return averageMapBykey
}

func sumArray(elems []float64) float64 {
	var sum float64
	for _, value := range elems {
		sum += value
	}
	return sum
}
