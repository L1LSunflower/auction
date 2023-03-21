package auctions

type AmountOffer struct {
	Amount float64 `json:"amount" validate:"required"`
}
