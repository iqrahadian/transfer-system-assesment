package param

import (
	"github.com/iqrahadian/paperid-assesment/common"
)

type DisburseParam struct {
	TransactionID            string `validate:"required"`
	SenderID                 string `validate:"required"`
	SourceAccountID          string `validate:"required"`
	DestinationAccountID     common.NullString
	DestinationAccountNumber string  `validate:"required"`
	Amount                   float64 `validate:"required"`
}
