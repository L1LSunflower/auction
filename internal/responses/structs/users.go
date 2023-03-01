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
