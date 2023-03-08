package entities

import "time"

const (
	CreditType = "credit"
	DebitType  = "debit"
)

type Transaction struct {
	ID        string
	Amount    float64
	Type      string
	CreatedAt time.Time
}
