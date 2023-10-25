package param

type DisburseParam struct {
	TransactionID            string  `validate:"required"`
	SenderID                 string  `validate:"required"`
	SourceAccountID          string  `validate:"required"`
	DestinationAccountNumber string  `validate:"required"`
	Amount                   float64 `validate:"required"`
}
