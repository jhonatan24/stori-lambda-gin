package dto

import (
	"fmt"
	"strconv"
	"time"
)

type TransactionFile struct {
	Id     string `csv:"id"`
	Date   string `csv:"date"`
	Amount string `csv:"transaction"`
}

func (t TransactionFile) GetDateFormat() time.Time {
	currentYear := time.Now().Year()
	strWithYear := fmt.Sprintf("%d/%s", currentYear, t.Date)
	layout := "2006/1/2"
	time, err := time.Parse(layout, strWithYear)
	if err != nil {
		fmt.Println("Error converting string to time.Time:", err)
	}
	return time
}

func (t TransactionFile) GetMoth() time.Month {
	return t.GetDateFormat().Month()
}

func (t TransactionFile) GetAmountFloat() float64 {
	amount, _ := strconv.ParseFloat(t.Amount, 64)
	return amount
}

func (t TransactionFile) IsDebit() bool {
	return t.GetAmountFloat() < 0
}
