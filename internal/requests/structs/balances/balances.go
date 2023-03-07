package balances

type Credit struct {
	ID     string
	Pan    string  `json:"pan"`
	CVC    string  `json:"cvc"`
	Amount float64 `json:"amount"`
}

type Debit struct {
	ID     string
	Amount float64 `json:"amount"`
}

type Balance struct {
	ID string
}
