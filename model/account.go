package model

type AccountType string

const (
	Internal AccountType = "Internal"
	BCA                  = "BCA"
	Mandiri              = "Mandiri"
)

type Account struct {
	ID            string
	UserID        string
	Name          string
	AccountNumber string
	Type          AccountType
	Balance       float64
}
