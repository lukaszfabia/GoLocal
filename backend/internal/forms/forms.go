package forms

import "time"

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

type EditAccount struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	Bio       string    `json:"bio"`
}

type RestoreAccount struct {
	Email string `json:"email"`
}

type CodeRequest struct {
	Code string `json:"code"`
}
