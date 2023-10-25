package account

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/model"
)

func GetAccountByID(carrier *ctx.Carrier, accountID string) (model.Account, common.Error)

func DeductAccountBalance(carrier *ctx.Carrier, accountID string, amount float64) common.Error

func UpdateAccountBalance(carrier *ctx.Carrier, accountID string, amount float64) common.Error
