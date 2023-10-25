package repo

import "github.com/iqrahadian/paperid-assesment/model"

var (
	UserRepo        = map[string]model.User{}
	AccountRepo     = map[string]model.Wallet{}
	TransactionRepo = map[string]model.Transaction{}
)

func OnLoad() {

	UserRepo["user1"] = model.User{
		ID:   "user1",
		Name: "Iqbal",
	}

	AccountRepo["account1"] = model.Wallet{
		ID:      "account1",
		Name:    "Iqbal's Account",
		UserID:  "user1",
		Balance: 100000,
	}

}
