package wallet

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/common/message"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/repo"
)

func GetAccountByID(carrier *ctx.Carrier, accountID string) (model.Wallet, common.Error) {

	if val, ok := repo.WalletRepo[accountID]; ok {
		return val, common.Error{}
	} else {
		return model.Wallet{}, common.Error{
			Error:   errors.New(message.DATA_NOT_FOUND),
			Message: message.DATA_NOT_FOUND,
			Code:    http.StatusNotFound,
		}
	}

}

func DeductAccountBalance(carrier *ctx.Carrier, accountID string, amount float64) common.Error {

	accountWallet, err := GetAccountByID(carrier, accountID)
	if err.Error != nil {
		return err
	}

	accountWallet.Balance -= amount
	repo.WalletRepo[accountID] = accountWallet

	fmt.Println("New Account Wallet : ", accountWallet)

	return common.Error{}

}

func IncreaseAccountBalance(carrier *ctx.Carrier, accountID string, amount float64) common.Error {

	accountWallet, err := GetAccountByID(carrier, accountID)
	if err.Error != nil {
		return err
	}

	accountWallet.Balance += amount
	repo.WalletRepo[accountID] = accountWallet

	return common.Error{}

}
