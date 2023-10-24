package param

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/model"
)

type DisburseParam struct {
	SourceAccountID          string `validate:"required"`
	DestinationAccountID     common.NullString
	DestinationAccountType   model.AccountType `validate:"required"`
	DestinationAccountNumber string            `validate:"required"`
	Amount                   float64           `validate:"required"`
}
