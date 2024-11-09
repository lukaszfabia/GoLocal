package forms

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RefreshTokenRequest struct {
	Token string `json:"refresh"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
