package repo

import "github.com/iqrahadian/paperid-assesment/model"

var (
	UserRepo        = map[string]model.User{}
	WalletRepo      = map[string]model.Wallet{}
	TransactionRepo = map[string]model.Transaction{}
)

func OnLoad() {

	UserRepo["user1"] = model.User{
		ID:   "user1",
		Name: "Iqbal",
	}

	WalletRepo["wallet1"] = model.Wallet{
		ID:      "wallet1",
		Name:    "Iqbal's Account",
		UserID:  "user1",
		Balance: 100000,
	}

}
