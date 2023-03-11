package structs

type SignUp struct {
	ID    string `json:"id,required"`
	Phone string `json:"phone,required"`
	Date  string `json:"date,required"`
}

type UserToken struct {
	ID           string `json:"id,required"`
	AccessToken  string `json:"access,required"`
	RefreshToken string `json:"refresh,required"`
}

type User struct {
	ID        string `json:"id,required"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
}

type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type OtpSent struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UpdateUser struct {
	Status    string `json:"status"`
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
}

type DeleteUser struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

type Profile struct {
	Status   string            `json:"status"`
	Balance  float64           `json:"balance"`
	Auctions []AuctionWithFile `json:"auctions"`
}

type ProfileHistory struct {
	Status   string            `json:"status"`
	Auctions []AuctionWithFile `json:"auctions"`
}
