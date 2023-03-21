package users

type SignUp struct {
	Phone     string `json:"phone" validate:"required,e164"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	City      string `json:"city" validate:"omitempty"`
}

type SignIn struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Tokens struct {
	ID           string `validate:"required"`
	AccessToken  string `json:"access" validate:"required"`
	RefreshToken string `json:"refresh" validate:"required"`
}

type Confirm struct {
	ID    string `validate:"required"`
	Phone string `json:"phone" validate:"required,e164"`
	Code  string `json:"code" validate:"required,len=4"`
}

type User struct {
	ID string `validate:"required"`
}

type Update struct {
	ID        string `validate:"required"`
	Email     string `json:"email" validate:"omitempty,email"`
	FirstName string `json:"first_name" validate:"omitempty"`
	LastName  string `json:"last_name" validate:"omitempty"`
	Password  string `json:"password" validate:"omitempty"`
}

type Delete struct {
	ID string `validate:"required"`
}

type RestorePassword struct {
	Phone string `json:"phone" validate:"required,phone"`
}

type ChangePassword struct {
	Phone    string `json:"phone" validate:"required,phone"`
	Code     string `json:"code" validate:"required,len=4"`
	Password string `json:"password" validate:"required"`
}

type AuthWS struct {
	ID     string `json:"id" validate:"required,uuid"`
	Access string `json:"access" validate:"required"`
}
