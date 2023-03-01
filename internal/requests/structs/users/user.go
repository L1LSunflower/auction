package users

type SignUp struct {
	Phone     string `json:"phone,required"`
	FirstName string `json:"first_name,required"`
	LastName  string `json:"last_name,required"`
	Email     string `json:"email,required"`
	Password  string `json:"password,required"`
	City      string `json:"city,omitempty"`
}

type SignIn struct {
	Phone    string `json:"phone,required"`
	Password string `json:"password,required"`
}

type Tokens struct {
	ID           string
	AccessToken  string
	RefreshToken string `json:"refresh_token"`
}

type Confirm struct {
	ID    string
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type User struct {
	ID string
}

type Update struct {
	ID        string
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type Delete struct {
	ID string
}

type RestorePassword struct {
	Phone string `json:"phone,required"`
}

type RefreshPassword struct {
	AccessToken  string
	RefreshToken string
}

type ChangePassword struct {
	Phone    string `json:"phone"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
