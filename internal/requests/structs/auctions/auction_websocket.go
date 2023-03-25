package auctions

type WSAuth struct {
	Bearer string `json:"bearer" validate:"required"`
}

type AmountOffer struct {
	Amount float64 `json:"amount" validate:"required"`
}
