package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/common/http/response"
	"github.com/iqrahadian/paperid-assesment/common/message"
	"github.com/iqrahadian/paperid-assesment/domain/transaction"
	"github.com/iqrahadian/paperid-assesment/domain/wallet"
	"github.com/iqrahadian/paperid-assesment/event"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func Disburse(res http.ResponseWriter, req *http.Request) {

	userID := req.Header.Get("user_id")
	carrier := ctx.Carrier{
		Ctx:    req.Context(),
		UserId: &userID,
	}

	var disburseParam param.DisburseParam

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&disburseParam); err != nil {
		response.Error(res, common.Error{err, http.StatusBadRequest, message.BAD_REQUEST})
		return
	}

	// validate account is exist & owned by use
	userWallet, err := wallet.GetAccountByID(&carrier, disburseParam.SourceAccountID)
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	// first check of balance to be used
	// this mean we can accept the transfer request
	// there will be another validation in case the state already changed
	if disburseParam.Amount > userWallet.Balance {
		err = common.Error{
			Message: message.INSUFFICIENT_BALANCE,
			Error:   errors.New(message.INSUFFICIENT_BALANCE),
			Code:    http.StatusBadRequest,
		}
		response.Error(res, err)
		return
	}

	transaction, err := transaction.CreatePendingTransaction(&carrier, disburseParam)
	if err.Error != nil {
		response.Error(res, err)
		return
	}
	disburseParam.TransactionID = transaction.ID

	// submit disburse request to event queue
	event.PublishEvent(disburseParam)

	response.Success(res, message.SUCCESS, http.StatusOK, nil)
}
