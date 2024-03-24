package domain

import (
	"github.com/stretchr/testify/assert"
	"stori-lambda/internal/dto"
	"testing"
)

func Test_file_with_once_element(t *testing.T) {
	key := "September"
	row := []dto.TransactionFile{
		{Amount: "+32", Date: "9/23", Id: "1"},
	}
	transactionReport := NewTransactionReport(row)
	assert.NotNil(t, transactionReport.CountTransactionByMouth)
	assert.NotZerof(t, transactionReport.TotalAmount, "vacio total amount")
	assert.Equal(t, 1, transactionReport.CountTransactionByMouth[key])
	assert.Equal(t, float64(0), transactionReport.AverageDebitAmountByMouth[key])
	assert.Equal(t, float64(32), transactionReport.AverageCreditAmountByMouth[key])
}

func Test_when_exits_same_moth(t *testing.T) {
	key := "September"
	row := []dto.TransactionFile{
		{Amount: "+32", Date: "9/23", Id: "1"},
		{Amount: "+32", Date: "9/23", Id: "2"},
	}
	transactionReport := NewTransactionReport(row)
	assert.NotNil(t, transactionReport.CountTransactionByMouth)
	assert.NotZerof(t, transactionReport.TotalAmount, "vacio total amount")
	assert.Equal(t, 2, transactionReport.CountTransactionByMouth[key])
	assert.Equal(t, float64(0), transactionReport.AverageDebitAmountByMouth[key])
	assert.Equal(t, float64(32), transactionReport.AverageCreditAmountByMouth[key])
}
