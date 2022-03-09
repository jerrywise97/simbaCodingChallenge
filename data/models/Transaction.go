package models

import (
	"github.com/jinzhu/gorm"
)

type TransactionStatus string
type TransactionType string

const (
	SUCCESS  TransactionStatus = "success"
	FAILED   TransactionStatus = "failed"
	TRANSFER TransactionType   = "transfer"
	CONVERT  TransactionType   = "convert"
)

type Transaction struct {
	gorm.Model
	Amount          float64
	Currency        Currency
	UserID          uint
	ReceiversID     uint
	TransactionType TransactionType
	Status          TransactionStatus
}