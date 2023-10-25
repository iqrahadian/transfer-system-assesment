package transaction

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func Disburse(carrier *ctx.Carrier, disburseParam param.DisburseParam) common.Error

func UpdateTransactionFailed(carrier *ctx.Carrier, transactionID string, status model.TransactionStatus, reason string) common.Error

func UpdateTransactionSuccess(carrier *ctx.Carrier, transactionID string, status model.TransactionStatus) common.Error
