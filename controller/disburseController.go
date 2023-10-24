package controller

import (
	"errors"
	"net/http"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/common/http/response"
	"github.com/iqrahadian/paperid-assesment/common/message"
	controllervalidator "github.com/iqrahadian/paperid-assesment/controller/controller_validator"
	"github.com/iqrahadian/paperid-assesment/domain/account"
)

func Disburse(res http.ResponseWriter, req *http.Request) {

	userID := req.Header.Get("user_id")
	carrier := ctx.Carrier{
		Ctx:    req.Context(),
		UserId: &userID,
	}

	disburseParam, err := controllervalidator.PostDirburseValidator(&carrier, req)
	if err.Error != nil {
		response.Error(res, err)
	}

	// validate account is exist & owned by use
	accountWallet, err := account.GetAccountByID(&carrier, "")
	if err.Error != nil {
		response.Error(res, err)
	}

	// first check of balance to be used
	// this mean we can accept the transfer request
	// there will be another validation in case the state already changed
	if disburseParam.Amount > accountWallet.Balance {
		err = common.Error{
			Message: message.INSUFFICIENT_BALANCE,
			Error:   errors.New(message.INSUFFICIENT_BALANCE),
			Code:    http.StatusBadRequest,
		}
		response.Error(res, err)
	}

	response.Success(res, message.SUCCESS, http.StatusOK, nil)
}
