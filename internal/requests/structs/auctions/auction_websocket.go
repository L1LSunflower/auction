package auctions

type WSAuth struct {
	Bearer string `json:"bearer" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type WSClose struct {
	Close bool `json:"close"`
}

type AmountOffer struct {
	Amount float64 `json:"amount" validate:"required"`
}
