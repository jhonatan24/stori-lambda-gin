package entity

import (
	"time"
)

type TransactionDetailsEntity struct {
	Id       uint `gorm:"type:int;primary_key;"`
	Date     time.Time
	Amount   float64   `gorm:"type:int;"`
	CreateAt time.Time `gorm:"autoCreateTime"`
}

func (t TransactionDetailsEntity) IsDebit() bool {
	return t.Amount < 0
}
