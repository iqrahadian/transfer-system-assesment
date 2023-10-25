package event

import (
	"context"
	"fmt"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/domain/disburse"
	"github.com/iqrahadian/paperid-assesment/domain/transaction"
	"github.com/iqrahadian/paperid-assesment/domain/wallet"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/model/param"
	"github.com/iqrahadian/paperid-assesment/repo"
)

func executeDisburse(job param.DisburseParam) {

	defer UnlockAccount(job.SourceAccountID)

	var err common.Error
	carrier := ctx.Carrier{
		Ctx:    context.Background(),
		UserId: &job.SenderID,
	}

	// validate amount
	userWallet, err := wallet.GetAccountByID(&carrier, job.SourceAccountID)
	if err.Error != nil {
		err = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
		if err.Error != nil {
			fmt.Errorf("Failed to update transaction status, cause : ", err.Error)
		}
	}

	// deduct wallet
	if userWallet.Balance > job.Amount {
		err = wallet.DeductAccountBalance(&carrier, job.SenderID, job.Amount)
		if err.Error != nil {
			fmt.Errorf("Failed to deduct account balance, cause : ", err.Error)

			_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, "Failed to deduct account balance")
		}
	}

	// after wallet balance deducted, open the lock allowing another transaction to be processed
	UnlockAccount(job.SourceAccountID)

	// try to disburse
	disburseProcessor, err := disburse.GetDisburseProcessor(job.DestinationAccountNumber)
	if err.Error != nil {
		fmt.Errorf("Failed to deduct account balance, cause : ", err.Error)
		_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
	}

	err = disburseProcessor.Disburse(job)
	if err.Error != nil {
		fmt.Errorf("Failed to disburse money, cause : ", err.Error)
		_ = transaction.UpdateTransactionFailed(&carrier, job.TransactionID, model.TransansactionFailed, err.Message)
		_ = wallet.IncreaseAccountBalance(&carrier, job.SourceAccountID, job.Amount)
	}

	// handle success
	_ = transaction.UpdateTransactionSuccess(&carrier, job.TransactionID, model.TransansactionSuccess)

	fmt.Println("Finish disburse money")
	fmt.Println("TRANSACTION DATA : ", repo.TransactionRepo)
	fmt.Println("WALLET DATA : ", repo.AccountRepo)

}
