package structs

type Credit struct {
	Status  string  `json:"status"`
	ID      string  `json:"user_id"`
	Balance float64 `json:"balance"`
	Date    string  `json:"date"`
}

type Debit struct {
	Status  string  `json:"status"`
	ID      string  `json:"user_id"`
	Balance float64 `json:"balance"`
	Date    string  `json:"date"`
}

type Balance struct {
	Status  string  `json:"status"`
	ID      string  `json:"user_id"`
	Balance float64 `json:"balance"`
	Date    string  `json:"date"`
}
