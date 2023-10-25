package model

import "github.com/iqrahadian/paperid-assesment/common"

type TransactionType string

// this can be a lot of type depend on usecase
// INTERNAL is for transfer to account inside system
// EXTERNAL is for transfer to account not in system, such us to bank account
const (
	InternalTransfer TransactionType = "INTERNAL"
	ExternalTransfer                 = "EXTERNAL"
)

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
	DestinationAccountID     common.NullString
	DestinationAccountType   AccountType
	DestinationAccountNumber string
	Amount                   float64
	Cost                     float64
	Type                     TransactionType
	Status                   TransactionStatus
	FailureMessage           common.NullString
	RequestTime              common.NullDateTime
	ProcessTime              common.NullDateTime
	FinishedTime             common.NullDateTime
}
