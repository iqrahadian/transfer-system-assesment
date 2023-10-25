package transaction

import (
	"errors"
	"net/http"
	"time"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/common/message"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/model/param"
	"github.com/iqrahadian/paperid-assesment/repo"
)

func GetTransactionByID(carrier *ctx.Carrier, transactionId string) (model.Transaction, common.Error) {

	if val, ok := repo.TransactionRepo[transactionId]; ok {
		return val, common.Error{}
	} else {

		return model.Transaction{}, common.Error{
			Error:   errors.New(message.DATA_NOT_FOUND),
			Message: message.DATA_NOT_FOUND,
			Code:    http.StatusNotFound,
		}
	}

}

func CreatePendingTransaction(carrier *ctx.Carrier, disburseParam param.DisburseParam) (model.Transaction, common.Error) {

	t := time.Now()
	transactionID := t.Format("20060102150405")
	transaction := model.Transaction{
		ID:     transactionID,
		Status: model.TransactionPending,
	}

	repo.TransactionRepo[transactionID] = transaction

	return transaction, common.Error{}

}

func UpdateTransactionFailed(carrier *ctx.Carrier, transactionID string, status model.TransactionStatus, reason string) common.Error {

	transaction, err := GetTransactionByID(carrier, transactionID)
	if err.Error != nil {
		return err
	}

	transaction.Status = model.TransansactionFailed
	transaction.FailureMessage = common.ParseToNullString(reason)

	repo.TransactionRepo[transactionID] = transaction

	return common.Error{}

}

func UpdateTransactionSuccess(carrier *ctx.Carrier, transactionID string, status model.TransactionStatus) common.Error {

	transaction, err := GetTransactionByID(carrier, transactionID)
	if err.Error != nil {
		return err
	}

	transaction.Status = model.TransansactionSuccess

	repo.TransactionRepo[transactionID] = transaction

	return common.Error{}

}
