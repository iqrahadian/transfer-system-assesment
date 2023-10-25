package param

type DisburseParam struct {
	TransactionID            string  `validate:"required"`
	SenderID                 string  `json:"sender_id" validate:"required"`
	SourceAccountID          string  `json:"source_account_id" validate:"required"`
	DestinationAccountNumber string  `json:"destination_number" validate:"required"`
	Amount                   float64 `json:"amount" validate:"required"`
}
