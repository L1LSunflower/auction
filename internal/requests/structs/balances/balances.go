package balances

type Credit struct {
	ID     string  `validate:"required"`
	Pan    string  `json:"pan" validate:"required"`
	CVV    string  `json:"cvv" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type Debit struct {
	ID     string  `validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type Balance struct {
	ID string `validate:"required"`
}
