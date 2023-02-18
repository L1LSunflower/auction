package users

type SignUp struct {
	Phone     string `json:"phone,required"`
	FirstName string `json:"first_name,required"`
	LastName  string `json:"last_name,required"`
	Email     string `json:"email,required"`
	Password  string `json:"password,required"`
	City      int    `json:"city,omitempty"`
}

type SignIn struct {
	Phone    string `json:"phone,required"`
	Password string `json:"password,required"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string `json:"refresh_token"`
}

type Confirm struct {
	ID   string
	Code string `json:"code"`
}

type User struct {
	ID string
}

type RestorePassword struct {
	Phone       string `json:"phone,required"`
	OldPassword string `json:"old_password,required"`
	NewPassword string `json:"new_password,required"`
}
