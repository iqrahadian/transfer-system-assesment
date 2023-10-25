package model

import "github.com/iqrahadian/paperid-assesment/common"

type TransactionStatus string

const (
	TransactionPending      TransactionStatus = "PENDING"
	TransansactionProcessed                   = "PROCESSED"
	TransansactionSubmitted                   = "SUBMITTED"
	TransansactionCancelled                   = "CANCELLED"
	TransansactionFailed                      = "FAILED"
	TransansactionSuccess                     = "SUCCESS"
)

type Transaction struct {
	ID                       string
	UserID                   string
	SourceAccountID          string
	DestinationAccountNumber string
	Amount                   float64
	Status                   TransactionStatus
	FailureMessage           common.NullString
	RequestTime              common.NullDateTime
	ProcessTime              common.NullDateTime
	FinishedTime             common.NullDateTime
}
