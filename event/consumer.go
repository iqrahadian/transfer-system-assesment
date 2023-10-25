package event

import (
	"context"
	"fmt"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/domain/account"
	"github.com/iqrahadian/paperid-assesment/domain/disburse"
	"github.com/iqrahadian/paperid-assesment/domain/transaction"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func executeDisburse(job param.DisburseParam) {

	defer UnlockAccount(job.SourceAccountID)

	var error common.Error
	carrier := ctx.Carrier{
		Ctx:    context.Background(),
		UserId: &job.SenderID,
	}

	// validate amount
	wallet, err := account.GetAccountByID(&carrier, job.SourceAccountID)
	if err.Error != nil {
		err = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
		if err.Error != nil {
			fmt.Errorf("Failed to update transaction status, cause : ", err.Error)
		}
	}

	// deduct wallet
	if wallet.Balance > job.Amount {
		err = account.DeductAccountBalance(&carrier, job.SenderID, job.Amount)
		if err.Error != nil {
			fmt.Errorf("Failed to deduct account balance, cause : ", err.Error)

			_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, "Failed to deduct account balance")
		}
	}

	// after wallet balance deducted, open the lock allowing another transaction to be processed
	UnlockAccount(job.SourceAccountID)

	// try to disburse
	disburseProcessor, err := disburse.GetDisburseProcessor(wallet.Type, wallet.AccountNumber)
	if err.Error != nil {
		fmt.Errorf("Failed to deduct account balance, cause : ", err.Error)
		_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
	}

	err = disburseProcessor.Disburse(job)
	if err.Error != nil {
		fmt.Errorf("Failed to disburse money, cause : ", err.Error)
		_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
	}

	// handle success
	_ = transaction.UpdateTransactionSuccess(&carrier, job.TransactionID, model.TransansactionSuccess)

	fmt.Println(error)

}
